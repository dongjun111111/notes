package rpc

import (
	"errors"
	"galopush/internal/logs"
	"net/rpc"
	"sync"
	"time"
)

type RpcClient struct {
	id         string //用来标识RPC客户端
	serverAddr string
	mutex      sync.Mutex
	rpcClient  *rpc.Client
	status     chan int //1-success 0-failed
}

//NewClient return a *Client wicth contant a rpc.handle
func NewRpcClient(id, serverAddr string, stateChan chan int) (*RpcClient, error) {
	var client RpcClient
	client.id = id
	client.serverAddr = serverAddr
	client.status = stateChan

	var cli *rpc.Client
	var err error

	if cli, err = rpc.DialHTTP("tcp", serverAddr); err != nil {
		return &client, err
	}
	client.rpcClient = cli
	client.status <- 1
	return &client, err
}

func (p *RpcClient) connect() error {
	var err error
	var cli *rpc.Client
	cli, err = rpc.DialHTTP("tcp", p.serverAddr)
	if err == nil {
		p.rpcClient = cli
	}
	return err
}

func (p *RpcClient) ReConnect() error {
	//先关闭旧连接
	p.rpcClient.Close()

	var err error
	if err = p.connect(); err == nil {
		logs.Logger.Debug("connect to rpc server ", p.serverAddr, "success")
		p.status <- 1
		return nil
	}
	return err
}

func (p *RpcClient) Close() error {
	return p.rpcClient.Close()
}

//StartPing start heartbeat with the ticker 5 s
func (p *RpcClient) StartPing() {
	go func(client *RpcClient) {
		tk := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-tk.C:
				err := client.Ping()
				if err != nil {
					logs.Logger.Error("[rpc]ping rpc server err ", err, " server addr  ", client.serverAddr)
					p.status <- 0
					return
				}
			}
		}
	}(p)
}

//Register comet登记
func (p *RpcClient) Register(cometId, tcpAddr, wsAddr, rpcAddr string) error {
	var request CometRegister
	request.CometId = cometId
	request.TcpAddr = tcpAddr
	request.WsAddr = wsAddr
	request.RpcAddr = rpcAddr

	var respone Response

	if err := p.rpcClient.Call("RpcServer.Register", &request, &respone); err != nil {
		logs.Logger.Error("[rpc client] RpcServer.Register ", err)
		return err
	}
	if respone.Code != 0 {
		return ERROR_RESPONSE
	}
	return nil
}

//鉴权
func (p *RpcClient) Auth(id string, termtype int, code string) error {
	var request AuthRequest
	request.Id = id
	request.Termtype = termtype
	request.Code = code
	var respone Response

	if err := p.rpcClient.Call("RpcServer.Auth", &request, &respone); err != nil {
		logs.Logger.Error("[rpc client] RpcServer.Auth ", err)
		return err
	}
	if respone.Code != 0 {
		return ERROR_RESPONSE
	}
	return nil
}

//Notify 用户状态登记
func (p *RpcClient) Notify(id string, plat int, token string, state int, cometId string) error {
	logs.Logger.Debug("[rpcclient] report state id=", id, " plat=", plat, " token=", token, " state=", state, " comet=", cometId)
	var request StateNotify
	request.Id = id
	request.CometId = cometId
	request.Token = token
	request.Termtype = plat
	request.State = state

	var respone Response

	if err := p.rpcClient.Call("RpcServer.Notify", &request, &respone); err != nil {
		logs.Logger.Error("[rpc client] RpcServer.Notify ", err)
		return err
	}
	return nil
}

//MsgUpward 消息上行
func (p *RpcClient) MsgUpward(id string, termtype int, msg string) error {
	var request MsgUpwardRequst
	request.Id = id
	request.Termtype = termtype
	request.Msg = msg

	var respone Response

	if err := p.rpcClient.Call("RpcServer.MsgUpward", &request, &respone); err != nil {
		logs.Logger.Error("[rpc client] RpcServer.MsgUpward ", err)
		return err
	}
	if respone.Code != 0 {
		return ERROR_RESPONSE
	}
	return nil
}

//Kick 踢人下线
func (p *RpcClient) Kick(id string, plat int, token string, reason int) error {
	var request KickRequst
	request.Id = id
	request.Termtype = plat
	request.Token = token
	request.Reason = reason

	var respone Response

	err := p.rpcClient.Call("RpcServer.Kick", &request, &respone)
	if err != nil {
		logs.Logger.Error("[rpc client] RpcServer.Kick ", err)
		return err
	}
	return nil
}

//Push 下发消息
func (p *RpcClient) Push(msgType int, id string, termtype int, iosToken string, flag int, msg string) error {
	var request PushRequst
	request.Tp = msgType
	request.Flag = flag
	request.Id = id
	request.Termtype = termtype
	request.AppleToken = iosToken
	request.Msg = msg

	var respone Response

	err := p.rpcClient.Call("RpcServer.Push", &request, &respone)
	if err != nil {
		logs.Logger.Error("[rpc client] RpcServer.Push ", err)
		return err
	}
	if respone.Code != 0 {
		return errors.New("PUSH ERROR")
	}
	return nil
}

//Ping heartbeat
func (p *RpcClient) Ping() error {
	var request Ping
	var respone Response

	if err := p.rpcClient.Call("RpcServer.Ping", &request, &respone); err != nil {
		logs.Logger.Error("[rpc client] RpcServer.Ping ", err)
		return err
	}
	if respone.Code != RPC_RET_SUCCESS {
		return ERROR_RESPONSE
	}
	return nil
}
