/*
声明：本代码来自skoo87，淘宝工程师，非原创
用途：用来替代channel，提高海量数据收发的性能
缺点：暂时只支持一个写

案例：

ring := ringbuffer.NewRing(100, 1000)

// 一个写端
go func() {
    var wbuf *ringbuffer.Buffer

    for i := 0; i < 10000; i++ {
        wbuf = ring.Write(wbuf, i)
    }
    ring.Stop(wbuf)
}()

// 10个读端
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)

    go func() {
        defer wg.Done()

        var (
            rbuf *ringbuffer.Buffer
            e    interface{}
        )

        for {
            if e, rbuf = ring.Read(rbuf); rbuf == nil {
                break
            }
            log.Println(e.(int))
        }
    }()
}

wg.Wait()
*/

package ringbuffer

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Ring struct {
	buffers []*Buffer
	size    int32
	rpos    int32
	wpos    int32
	stop    int32
	firstW  *Buffer
}

func NewRing(rsize, bsize int) *Ring {
	r := new(Ring)
	r.buffers = make([]*Buffer, rsize)
	r.size = int32(rsize)
	r.rpos = 0
	r.wpos = 0
	r.stop = 0

	for i := 0; i < rsize; i++ {
		buf := newBuffer(bsize)
		buf.index = int32(i)
		buf.ring = r

		r.buffers[i] = buf
	}

	r.firstW = r.prepareWrite()

	return r
}

func (r *Ring) prepareRead() *Buffer {
	var rpos, wpos int32

	for i := 0; i < 10; i++ {
		rpos = atomic.LoadInt32(&r.rpos)
		wpos = atomic.LoadInt32(&r.wpos)

		if rpos == wpos {
			if atomic.LoadInt32(&r.stop) == 1 {
				return nil
			}
			runtime.Gosched()
			continue
		}

		break
	}

	buffer := r.buffers[rpos]
	buffer.prepareGet()
	return buffer
}

func (r *Ring) Read(buf *Buffer) (e interface{}, next *Buffer) {
	if buf == nil {
		if buf = r.prepareRead(); buf == nil {
			return nil, nil
		}
	}

	for {
		if e = buf.get(); e == nil {
			if buf = r.prepareRead(); buf == nil {
				return
			}
			continue
		}
		next = buf
		return
	}
}

func (r *Ring) ExitRead(buf *Buffer) {
	buf.exitGet()
}

func (r *Ring) prepareWrite() *Buffer {
	buffer := r.buffers[r.wpos]
	buffer.preparePut()
	return buffer
}

func (r *Ring) Write(buf *Buffer, e interface{}) *Buffer {
	if buf == nil {
		buf = r.firstW
	}

	for {
		if err := buf.put(e); err != nil {
			buf = r.NextWrite(buf)
			continue
		}
		return buf
	}
}

func (r *Ring) NextWrite(buf *Buffer) (next *Buffer) {
	return buf.nextPut()
}

func (r *Ring) Stop(buf *Buffer) {
	atomic.AddInt32(&r.stop, 1)
	buf.finishPut()
}

type Buffer struct {
	body     []interface{}
	w        int32
	r        int32
	lock     sync.RWMutex
	once     *sync.Once
	index    int32 // Buffer在Ring数组中的下标
	ring     *Ring
	readonly int32
}

func newBuffer(size int) *Buffer {
	buf := new(Buffer)
	buf.body = make([]interface{}, size)
	buf.once = new(sync.Once)
	return buf
}

func (b *Buffer) lastDo() {
	if b.index+1 == b.ring.size {
		atomic.StoreInt32(&b.ring.rpos, 0)
	} else {
		atomic.StoreInt32(&b.ring.rpos, b.index+1)
	}

	atomic.SwapInt32(&b.readonly, 0)
}

func (b *Buffer) prepareGet() {
	b.lock.RLock()
}

func (b *Buffer) finishGet() {
	b.once.Do(b.lastDo)
	b.lock.RUnlock()
}

func (b *Buffer) exitGet() {
	b.lock.RUnlock()
}

func (b *Buffer) get() (e interface{}) {
	defer func() {
		if e == nil {
			b.finishGet()
		}
	}()

	if atomic.LoadInt32(&b.r) >= b.w {
		return
	}
	newr := atomic.AddInt32(&b.r, 1)
	if newr <= b.w {
		e = b.body[newr-1]
	}

	return
}

func (b *Buffer) preparePut() {
	for {
		if atomic.LoadInt32(&b.readonly) == 0 {
			break
		}
		// TODO fix sleep
		time.Sleep(time.Millisecond * time.Duration(b.ring.size) / 2)
	}
	b.lock.Lock()
	b.r = 0
	b.w = 0
	b.once = new(sync.Once)
}

func (b *Buffer) finishPut() {
	atomic.AddInt32(&b.ring.wpos, 1)
	atomic.CompareAndSwapInt32(&b.ring.wpos, b.ring.size, 0)

	atomic.SwapInt32(&b.readonly, 1)
	b.lock.Unlock()
}

func (b *Buffer) nextPut() *Buffer {
	atomic.AddInt32(&b.ring.wpos, 1)
	atomic.CompareAndSwapInt32(&b.ring.wpos, b.ring.size, 0)

	next := b.ring.buffers[b.ring.wpos]
	next.preparePut()

	atomic.SwapInt32(&b.readonly, 1)
	b.lock.Unlock()

	return next
}

func (b *Buffer) put(e interface{}) error {
	if int(b.w) >= len(b.body) {
		return errors.New("full")
	}

	b.body[b.w] = e
	b.w++

	return nil
}
