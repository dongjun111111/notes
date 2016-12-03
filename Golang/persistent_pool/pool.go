package pool

import (
	"errors"
	"sync"
	"time"
)

//持久化对象池配置
type PoolConfig struct {
	InitialCap     int
	MaxCap         int
	Factory        func() ([]byte, error)
	Close          func([]byte) error
	InitialTimeout time.Duration
}

type channelPool struct {
	mu            sync.Mutex
	values        chan *InitValue
	factory       func() ([]byte, error) //生成对象的方法
	close         func([]byte) error     //销毁对象的方法
	initalTimeout time.Duration
}

type InitValue struct {
	value []byte
	t     time.Time
}

func NewChannelPool(poolConfig *PoolConfig) (Pool, error) {
	if poolConfig.InitialCap < 0 || poolConfig.MaxCap <= 0 || poolConfig.InitialCap > poolConfig.MaxCap {
		return nil, errors.New("配置信息出错")
	}
	c := &channelPool{
		values:        make(chan *InitValue, poolConfig.MaxCap),
		factory:       poolConfig.Factory,
		close:         poolConfig.Close,
		initalTimeout: poolConfig.InitialTimeout,
	}

	for i := 0; i < poolConfig.InitialCap; i++ {
		value, err := c.factory()
		if err != nil {
			c.Release()
			return nil, errors.New("对象无法放入对象池中")
		}
		c.values <- &InitValue{value: value, t: time.Now()}
	}
	return c, nil
}

//获取所有对象
func (c *channelPool) GetValues() chan *InitValue {
	c.mu.Lock()
	values := c.values
	c.mu.Unlock()
	return values
}

func (c *channelPool) Get() ([]byte, error) {
	values := c.GetValues()
	if values == nil {
		return nil, ErrClosed
	}
	for {
		select {
		case wrapvalue := <-values:
			if wrapvalue == nil {
				return nil, ErrClosed
			}
			//超时则丢弃，强制置空
			if timeout := c.initalTimeout; timeout > 0 {
				if wrapvalue.t.Add(timeout).Before(time.Now()) {
					c.Close(wrapvalue.value)
					if wrapvalue.value != nil {
						wrapvalue.value = nil
					}
					continue
				}
			}
			return wrapvalue.value, nil
		default:
			value, err := c.factory()
			if err != nil {
				return nil, err
			}
			return value, err
		}
	}
}

//将对象放回对象池中
func (c *channelPool) Put(value []byte) error {
	if value == nil {
		return errors.New("对象不可用")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.values == nil {
		return c.Close(value)
	}
	select {
	case c.values <- &InitValue{value: value, t: time.Now()}:
		return nil
	default:
		return c.Close(value)
	}
}

func (c *channelPool) Close(value []byte) error {
	if value == nil {
		return errors.New("空对象警告")
	}
	value = nil
	return nil
}

//释放对象池
func (c *channelPool) Release() {
	c.mu.Lock()
	values := c.values
	c.values = nil
	c.factory = nil
	closeFunc := c.close
	c.close = nil
	c.mu.Unlock()
	if values == nil {
		return
	}
	close(values)
	for wrapvalues := range values {
		closeFunc(wrapvalues.value)
	}
}

//获取长度
func (c *channelPool) Len() int {
	return len(c.GetValues())
}
