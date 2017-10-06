package rds

import (
	"encoding/json"
	"galopush/internal/logs"
	"sync"

	"github.com/hoisie/redis"
)

//存储器规则
//查找：优先查找内存存储，再查询redis存储
//写入：优先写入redis存储，再写内存存储
type Storager struct {
	//redis存储
	cli *redis.Client

	//内存存储
	mutex    sync.Mutex
	memStore map[string]*Sessions
}

func NewStorager(dial, pswd string, db int) *Storager {
	var store Storager
	var client redis.Client
	client.Addr = dial
	client.Password = pswd
	client.Db = db
	client.MaxPoolSize = 10
	store.cli = &client
	store.memStore = make(map[string]*Sessions)
	return &store
}

//FindSessions 查找session组合 如果未找到则返回nil
func (p *Storager) FindSessions(id string) *Sessions {
	//内存拷贝出去 防止多线程操作失败
	var sess Sessions

	//先查询内存
	p.mutex.Lock()
	s, ok := p.memStore[id]
	p.mutex.Unlock()
	if ok {
		sess = *s
		return &sess
	}

	//如果没有找到则查询redis
	//var sess Sessions
	b, err := p.cli.Get(id)
	if err != nil {
		logs.Logger.Error("Redis err:", err)
		return nil
	}
	err = json.Unmarshal(b, &sess)
	if err != nil {
		logs.Logger.Error("Unmarshal err:", err, " b=", string(b))
		return nil
	}
	p.memStore[id] = &sess
	return &sess
}

//SaveSessions 保持session组合
func (p *Storager) SaveSessions(sess *Sessions) error {
	b, err := json.Marshal(sess)
	if err != nil {
		logs.Logger.Error("Marshal err:", err, " b=", string(b))
		return err
	}
	err = p.cli.Set(sess.Id, b)
	if err != nil {
		logs.Logger.Error("Redis err:", err)
		return err
	}
	p.mutex.Lock()
	p.memStore[sess.Id] = sess
	p.mutex.Unlock()
	return nil
}

//SessionOnline 返回指定用户是否在线
func (p *Storager) SessionOnline(id string, plat int) bool {
	var online bool
	sess := p.FindSessions(id)
	if sess == nil {
		return online
	}
	for _, v := range sess.Sess {
		if v.Plat == plat && v.Online == true {
			online = true
			return online
		}
	}
	return online
}

//SessionComet 返回指定用户所连接comet
func (p *Storager) SessionComet(id string) string {
	sess := p.FindSessions(id)
	if sess != nil {
		return sess.CometId
	}
	return ""
}

func (p *Storager) SessionCount() int {
	return len(p.memStore)
}

func (p *Storager) OfflineComet(comet string) {
	p.mutex.Lock()
	for _, sess := range p.memStore {
		if sess.CometId == comet {
			for _, it := range sess.Sess {
				it.Online = false
			}
			b, err := json.Marshal(sess)
			if err != nil {
				logs.Logger.Error("Marshal err:", err, " b=", string(b))
			}
			if err := p.cli.Set(sess.Id, b); err != nil {
				logs.Logger.Error("Redis err:", err)
			}
		}
	}
	p.mutex.Unlock()
}

/*
* 离线消息存取
 */
func (p *Storager) SavePushMsg(id string, plat int, msg []byte) error {
	logs.Logger.Debug("[Storager] SavePushMsg id=", id, " plat=", plat)
	sess := p.FindSessions(id)
	if sess == nil {
		return nil
	}

	for _, it := range sess.Sess {
		if it.Plat == plat {
			it.OffMsg.Push.OffCnt = it.OffMsg.Push.OffCnt + 1
			it.OffMsg.Push.Msg = make([]byte, 0)
			it.OffMsg.Push.Msg = append(it.OffMsg.Push.Msg, msg...)
			break
		}
	}
	p.SaveSessions(sess)
	return nil
}

func (p *Storager) SaveCallbackMsg(id string, plat int, msg []byte) error {
	logs.Logger.Debug("[Storager] SaveCallbackMsg id=", id, " plat=", plat)
	sess := p.FindSessions(id)
	if sess == nil {
		return nil
	}

	for _, it := range sess.Sess {
		if it.Plat == plat {
			for _, v := range it.OffMsg.CB {
				if string(v.Msg) == string(msg) {
					return nil
				}
			}
			var i Item
			i.Msg = append(i.Msg, msg...)
			it.OffMsg.CB = append(it.OffMsg.CB, &i)
			break
		}
	}
	p.SaveSessions(sess)
	return nil
}

func (p *Storager) SaveImMsg(id string, plat int, msg []byte) error {
	logs.Logger.Debug("[Storager] SaveImMsg id=", id, " plat=", plat)
	sess := p.FindSessions(id)
	if sess == nil {
		return nil
	}

	var bWebOnline bool
	for _, it := range sess.Sess {
		if it.Plat == 8 {
			if it.Online == true {
				bWebOnline = true
			}
			break
		}
	}

	for _, it := range sess.Sess {
		if it.Plat == plat {
			if len(it.OffMsg.Im) > 50 {
				it.OffMsg.Im = it.OffMsg.Im[1:50]
			}

			var i Item
			i.Msg = append(i.Msg, msg...)
			i.WebOnline = bWebOnline
			it.OffMsg.Im = append(it.OffMsg.Im, &i)
			break
		}
	}
	p.SaveSessions(sess)
	return nil
}

//实际上目前只有android有存储离线push web不存 ios走apns
func (p *Storager) GetPushMsg(id string, plat int) *Push {
	logs.Logger.Debug("[Storager] GetPushMsg id=", id, " plat=", plat)
	var r *Push
	sess := p.FindSessions(id)
	if sess == nil {
		return r
	}
	for _, it := range sess.Sess {
		if it.Plat == plat {
			r = &it.OffMsg.Push
			it.OffMsg.Push.OffCnt = 0
			it.OffMsg.Push.Msg = make([]byte, 0)
			break
		}
	}

	p.SaveSessions(sess)
	return r
}

func (p *Storager) GetCallbackMsg(id string, plat int) []*Item {
	logs.Logger.Debug("[Storager] GetCallbackMsg id=", id, " plat=", plat)
	var r []*Item
	sess := p.FindSessions(id)
	if sess == nil {
		return r
	}
	for _, it := range sess.Sess {
		if it.Plat == plat {
			r = it.OffMsg.CB
			it.OffMsg.CB = make([]*Item, 0)
			break
		}
	}

	p.SaveSessions(sess)
	return r
}

func (p *Storager) GetImMsg(id string, plat int) []*Item {
	logs.Logger.Debug("[Storager] GetImMsg id=", id, " plat=", plat)
	var r []*Item
	sess := p.FindSessions(id)
	if sess == nil {
		return r
	}
	for _, it := range sess.Sess {
		if it.Plat == plat {
			r = it.OffMsg.Im
			it.OffMsg.Im = make([]*Item, 0)
			break
		}
	}

	p.SaveSessions(sess)
	return r
}
