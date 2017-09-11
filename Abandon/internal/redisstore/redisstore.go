// redisstore 实现离线消息的redis存储与获取
package redisstore

import (
	"encoding/json"
	"galopush/internal/logs"

	"github.com/hoisie/redis"
)

type OffMsg struct {
	Id   string `json:"id"`
	Body Body   `json:"body"`
}

type Body struct {
	Push     Push   `json:"push"`
	Callback []Item `json:"callback"`
	Im       []Item `json:"im"`
}

type Push struct {
	Offline uint16 `json:"offline"`
	Msg     []byte `json:"msg"`
}

type Item struct {
	Plat      int    `json:"plat"`
	WebOnline int    `json:"webOnline"`
	Msg       []byte `json:"msg"`
}

type Storager struct {
	cli *redis.Client
}

func NewStorager(dial, pswd string, db int) *Storager {
	var store Storager
	var client redis.Client
	client.Addr = dial
	client.Password = pswd
	client.Db = db
	client.MaxPoolSize = 10
	store.cli = &client
	return &store
}

//func (p *Storager) GetOfflineMsg(id string) *OffMsg {
//	var omsg OffMsg
//	body, err := p.cli.Get(id)
//	if err != nil {
//		logs.Logger.Error(err)
//		return nil
//	}
//	json.Unmarshal(body, &omsg)
//	return &omsg
//}

func (p *Storager) SavePushMsg(id string, msg []byte) error {
	logs.Logger.Debug("[Storager] SavePushMsg id=", id, " msg=", string(msg[:]))
	b, err := p.cli.Exists(id)
	if err != nil {
		logs.Logger.Error(err)
		return err
	}
	var omsg OffMsg
	switch b {
	case true:
		body, err := p.cli.Get(id)
		if err != nil {
			logs.Logger.Error(err)
			return err
		}
		json.Unmarshal(body, &omsg)
		omsg.Body.Push.Offline = omsg.Body.Push.Offline + 1
		omsg.Body.Push.Msg = make([]byte, 0)
		omsg.Body.Push.Msg = msg
	case false:
		omsg.Body.Push.Offline = omsg.Body.Push.Offline + 1
		omsg.Body.Push.Msg = make([]byte, 0)
		omsg.Body.Push.Msg = msg
	}
	by, _ := json.Marshal(&omsg)
	err = p.cli.Set(id, by)
	return err
}

func (p *Storager) SaveCallbackMsg(id string, plat int, msg []byte) error {
	logs.Logger.Debug("[Storager] SaveCallbackMsg id=", id, " plat=", plat, " msg=", string(msg[:]))
	b, err := p.cli.Exists(id)
	if err != nil {
		logs.Logger.Error(err)
		return err
	}
	var omsg OffMsg
	switch b {
	case true:
		body, err := p.cli.Get(id)
		if err != nil {
			logs.Logger.Error(err)
			return err
		}
		json.Unmarshal(body, &omsg)

		//校验是否有重复回调 added by ligang @ 2016-08-01
		for _, v := range omsg.Body.Callback {
			if string(v.Msg) == string(msg) {
				break
			}
		}

		var it Item
		it.Msg = append(it.Msg, msg...)
		it.Plat = plat
		omsg.Body.Callback = append(omsg.Body.Callback, it)
	case false:
		var it Item
		it.Msg = append(it.Msg, msg...)
		it.Plat = plat
		omsg.Body.Callback = append(omsg.Body.Callback, it)
	}
	by, _ := json.Marshal(&omsg)
	if err = p.cli.Set(id, by); err != nil {
		logs.Logger.Error(err)
	}
	return err
}

func (p *Storager) SaveImMsg(id string, plat, webOnline int, msg []byte) error {
	logs.Logger.Debug("[Storager] SaveImMsg id=", id, " plat=", plat, " msg=", string(msg[:]))
	b, err := p.cli.Exists(id)
	if err != nil {
		logs.Logger.Error(err)
		return err
	}
	var omsg OffMsg
	switch b {
	case true:
		body, err := p.cli.Get(id)
		if err != nil {
			logs.Logger.Error(err)
			return err
		}
		json.Unmarshal(body, &omsg)
		//只存50条
		if len(omsg.Body.Im) > 50 {
			omsg.Body.Im = omsg.Body.Im[1:50]
		}
		var it Item
		it.Msg = append(it.Msg, msg...)
		it.Plat = plat
		it.WebOnline = webOnline
		omsg.Body.Im = append(omsg.Body.Im, it)
	case false:
		var it Item
		it.Msg = append(it.Msg, msg...)
		it.Plat = plat
		it.WebOnline = webOnline
		omsg.Body.Im = append(omsg.Body.Im, it)
	}
	by, _ := json.Marshal(&omsg)
	if err = p.cli.Set(id, by); err != nil {
		logs.Logger.Error(err)
	}
	return err
}

func (p *Storager) GetPushMsg(id string) (uint16, []byte) {
	var offCount uint16
	var buff []byte

	var omsg OffMsg
	body, err := p.cli.Get(id)
	if err != nil {
		logs.Logger.Error(err)
		return offCount, buff
	}
	json.Unmarshal(body, &omsg)
	offCount = omsg.Body.Push.Offline
	buff = append(buff, omsg.Body.Push.Msg...)

	//重置消息
	omsg.Body.Push.Offline = 0
	omsg.Body.Push.Msg = make([]byte, 0)
	by, _ := json.Marshal(&omsg)
	if err = p.cli.Set(id, by); err != nil {
		logs.Logger.Error(err)
	}
	return offCount, buff
}

func (p *Storager) GetCallbackMsg(id string, plat int) [][]byte {
	var buff [][]byte

	var omsg OffMsg
	body, err := p.cli.Get(id)
	if err != nil {
		logs.Logger.Error(err)
		return buff
	}
	json.Unmarshal(body, &omsg)
	var newItem []Item
	for _, v := range omsg.Body.Callback {
		var it Item
		it = v
		if it.Plat == plat {
			buff = append(buff, it.Msg)
		} else {
			newItem = append(newItem, it)
		}
	}

	//重置消息
	omsg.Body.Callback = newItem
	by, _ := json.Marshal(&omsg)
	if err = p.cli.Set(id, by); err != nil {
		logs.Logger.Error(err)
	}
	return buff
}

func (p *Storager) GetImMsg(id string, plat int) []Item {
	var rItem []Item
	var newItem []Item
	var omsg OffMsg
	body, err := p.cli.Get(id)
	if err != nil {
		logs.Logger.Error(err)
		return rItem
	}
	json.Unmarshal(body, &omsg)
	for _, v := range omsg.Body.Im {
		var it Item
		it = v
		if it.Plat == plat {
			rItem = append(rItem, it)
		} else {
			newItem = append(newItem, it)
		}
	}

	//重置消息
	omsg.Body.Im = newItem
	by, _ := json.Marshal(&omsg)
	if err = p.cli.Set(id, by); err != nil {
		logs.Logger.Error(err)
	}
	return rItem
}
