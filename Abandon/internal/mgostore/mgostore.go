// mgostore 实现离线消息的mongoDB存储与获取
package mgostore

import (
	"galopush/logs"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type offMsg struct {
	Id   string `json:"id"`
	Body body   `json:"body"`
}

type body struct {
	Push     push   `json:"push"`
	Callback []item `json:"callback"`
	Im       []item `json:"im"`
}

type push struct {
	Offline uint16 `json:"offline"`
	Msg     []byte `json:"msg"`
}

type item struct {
	Plat int    `json:"plat"`
	Msg  []byte `json:"msg"`
}

type Storager struct {
	session    *mgo.Session
	mdb        *mgo.Database
	collection *mgo.Collection
}

func NewStorager(dial, db, collection string) *Storager {
	var store Storager
	session, err := mgo.Dial(dial)
	if err != nil {
		logs.Logger.Error("[mongodb]mgo.Dial ", err)
		return nil
	}
	session.SetMode(mgo.Monotonic, true)
	store.session = session
	store.mdb = session.DB(db)
	store.collection = store.mdb.C(collection)
	return &store
}

func (p *Storager) GetMsg(id string) (*offMsg, error) {
	var msg offMsg
	err := p.collection.Find(bson.M{"id": id}).One(&msg)
	if err != nil {
		msg.Id = id
		err := p.collection.Insert(&msg)
		if err != nil {
			logs.Logger.Error("[mongodb]p.collection.Insert ", err)
			return &msg, err
		}
	}
	return &msg, nil
}

func (p *Storager) SavePushMsg(id string, msg []byte) error {
	logs.Logger.Debug("[mstore] SavePushMsg id=", id, " msg=", string(msg[:]))
	m, err := p.GetMsg(id)
	if err != nil {
		return err
	}
	m.Body.Push.Offline = m.Body.Push.Offline + 1
	m.Body.Push.Msg = make([]byte, 0)
	m.Body.Push.Msg = msg
	return p.collection.Update(bson.M{"id": id}, &m)
}

func (p *Storager) SaveCallbackMsg(id string, plat int, msg []byte) error {
	logs.Logger.Debug("[mstore] SaveCallbackMsg id=", id, " plat=", plat, " msg=", string(msg[:]))
	m, err := p.GetMsg(id)
	if err != nil {
		return err
	}
	var it item
	it.Msg = append(it.Msg, msg...)
	it.Plat = plat
	m.Body.Callback = append(m.Body.Callback, it)
	return p.collection.Update(bson.M{"id": id}, &m)
}

func (p *Storager) SaveImMsg(id string, plat int, msg []byte) error {
	logs.Logger.Debug("[mstore] SaveImMsg id=", id, " plat=", plat, " msg=", string(msg[:]))
	m, err := p.GetMsg(id)
	if err != nil {
		return err
	}
	var it item
	it.Msg = append(it.Msg, msg...)
	it.Plat = plat
	m.Body.Im = append(m.Body.Im, it)
	return p.collection.Update(bson.M{"id": id}, &m)
}

func (p *Storager) GetPushMsg(id string) (uint16, []byte) {
	var offCount uint16
	var msg []byte
	m, err := p.GetMsg(id)
	if err != nil {
		return offCount, msg
	}

	offCount = m.Body.Push.Offline
	msg = m.Body.Push.Msg
	m.Body.Push.Offline = 0
	m.Body.Push.Msg = make([]byte, 0)
	p.collection.Update(bson.M{"id": id}, &m)
	return offCount, msg
}

func (p *Storager) GetCallbackMsg(id string, plat int) [][]byte {
	var buff [][]byte
	m, err := p.GetMsg(id)
	if err != nil {
		return buff
	}
	var newItem []item
	for _, v := range m.Body.Callback {
		var it item
		it = v
		if v.Plat == plat {
			buff = append(buff, it.Msg)
		} else {
			newItem = append(newItem, it)
		}
	}
	m.Body.Callback = newItem
	p.collection.Update(bson.M{"id": id}, &m)
	return buff
}

func (p *Storager) GetImMsg(id string, plat int) [][]byte {
	var buff [][]byte
	m, err := p.GetMsg(id)
	if err != nil {
		return buff
	}
	var newItem []item
	for _, v := range m.Body.Im {
		var it item
		it = v
		if v.Plat == plat {
			buff = append(buff, it.Msg)
		} else {
			newItem = append(newItem, it)
		}
	}
	m.Body.Im = newItem
	p.collection.Update(bson.M{"id": id}, &m)
	return buff
}
