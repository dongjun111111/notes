package gobconn

import (
	"encoding/gob"
	"errors"
	"net"
	"reflect"
	"sync"
	"unsafe"
)

type message struct {
	Type  string
	value reflect.Value
}

func (self message) Recovery() {
	putPointer(self.value)
	putMsg(self)
}

func (self message) Interface() interface{} {
	return self.value.Elem().Interface()
}

/*
声明一个消息池用来重用对象
*/

var msgPool sync.Pool

func getMsg() message {
	if msg, ok := msgPool.Get().(message); ok {
		return msg
	}
	return message{}
}

func putMsg(msg message) {
	msgPool.Put(msg)
}

type gobConnection struct {
	rwc   net.Conn
	enc   *gob.Encoder
	dec   *gob.Decoder
	rlock sync.Mutex
	wlock sync.Mutex
}

type GobConnection interface {
	Read() (msg message, err error)
	Write(msg interface{}) (err error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

var gobPool sync.Pool

func NewGobConnection(conn net.Conn) GobConnection {
	if gcn, ok := gobPool.Get().(*gobConnection); ok {
		gcn.rwc = conn
		gcn.enc = gob.NewEncoder(conn)
		gcn.dec = gob.NewDecoder(conn)
		return gcn
	}
	return &gobConnection{rwc: conn, enc: gob.NewEncoder(conn), dec: gob.NewDecoder(conn)}
}

type msgStruct struct {
	StructName string
}

var (
	rheadMsg = msgStruct{}
	wheadMsg = msgStruct{}
)

func (self *gobConnection) Read() (msg message, err error) {
	self.rlock.Lock()
	defer self.rlock.Unlock()

	err = self.dec.Decode(&rheadMsg)
	if err != nil {
		return
	}
	var typ reflect.Type
	typ, err = GetMsgType(rheadMsg.StructName)
	if err != nil {
		return
	}
	msg = getMsg()
	msg.Type = rheadMsg.StructName
	var value = getPointer(typ)
	err = self.dec.DecodeValue(value)
	if err != nil {
		msg.Recovery()
		return
	}
	msg.value = value
	return
}

func (self *gobConnection) Write(msg interface{}) (err error) {
	self.wlock.Lock()
	value := reflect.ValueOf(msg)
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Ptr {
		wheadMsg.StructName = value.Elem().Type().String()
	} else {
		wheadMsg.StructName = value.Type().String()
	}
	err = self.enc.Encode(wheadMsg)
	if err != nil {
		self.wlock.Unlock()
		return
	}
	err = self.enc.EncodeValue(value)
	self.wlock.Unlock()
	return
}

func (self *gobConnection) Close() error {
	self.enc = nil
	self.dec = nil
	err := self.rwc.Close()
	gobPool.Put(self)
	return err
}

func (self *gobConnection) LocalAddr() net.Addr {
	return self.rwc.LocalAddr()
}

func (self *gobConnection) RemoteAddr() net.Addr {
	return self.rwc.RemoteAddr()
}

/*
通过指定类型申请一个定长的内存.
*/

var (
	lock   sync.Mutex
	ptrMap = make(map[string]*sync.Pool)
)

func getPointer(typ reflect.Type) reflect.Value {
	p, ok := ptrMap[typ.String()]
	if ok {
		if value, ok := p.Get().(reflect.Value); ok {
			return value
		}
		return reflect.New(typ)
	}
	lock.Lock()
	ptrMap[typ.String()] = new(sync.Pool)
	lock.Unlock()
	return reflect.New(typ)
}

func putPointer(value reflect.Value) {
	elem := value.Elem().Type()
	p, ok := ptrMap[elem.String()]
	if !ok {
		lock.Lock()
		p = new(sync.Pool)
		ptrMap[elem.String()] = p
		lock.Unlock()
	}
	ClearData(elem.Size(), unsafe.Pointer(value.Pointer()))
	p.Put(value)
}

/*
 使用此包进行数据发送之前必须将类型注册.否则接收放无法解包
*/

var (
	typeMap   = make(map[string]reflect.Type)
	Errortype = errors.New("type not register")
)

func GetMsgType(name string) (reflect.Type, error) {
	typ, ok := typeMap[name]
	if ok {
		return typ, nil
	}
	return nil, Errortype
}

func GetMsgAllType() []string {
	list := make([]string, 0, len(typeMap))
	for name, _ := range typeMap {
		list = append(list, name)
	}
	return list
}

func RegisterType(typ reflect.Type) {
	typeMap[typ.String()] = typ
}

func DeleteType(name string) {
	delete(typeMap, name)
}

/*
清除固定长度的内存数据,使用方法是:指定内存开始地址,和长度.
请勿随便使用.使用不当可能会清除有效数据
*/

func ClearData(size uintptr, ptr unsafe.Pointer) {
	var temptr uintptr = uintptr(ptr)
	var step uintptr = 1
	for {
		if size <= 0 {
			break
		}
		switch {
		case 1 <= size && size < 8:
			step = 1
		case 8 <= size && size < 32:
			step = 8
		case 32 <= size && size < 64:
			step = 32
		case size >= 64:
			step = 64
		}
		clearData(step, unsafe.Pointer(temptr))
		temptr += step
		size -= step
	}
}

func clearData(size uintptr, ptr unsafe.Pointer) {
	switch size {
	case 1:
		*(*[1]byte)(ptr) = [1]byte{}
	case 8:
		*(*[8]byte)(ptr) = [8]byte{}
	case 32:
		*(*[32]byte)(ptr) = [32]byte{}
	case 64:
		*(*[64]byte)(ptr) = [64]byte{}
	}
}
