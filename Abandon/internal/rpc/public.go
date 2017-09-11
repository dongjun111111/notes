//rpc 提供comet to router，router to comet的RPC调用
package rpc

import (
	"errors"
)

var (
	ERROR_RESPONSE = errors.New("error response code")
)

const (
	RPC_RET_SUCCESS = 0
	RPC_RET_FAILED  = -1
)

const (
	MSG_TYPE_UNKNOW   = 0
	MSG_TYPE_PUSH     = 1
	MSG_TYPE_CALLBACK = 2
	MSG_TYPE_IM       = 3
)

const (
	STATE_OFFLINE = 0
	STATE_ONLINE  = 1
)

/* Comet向Router注册
* 方向:Comet->Router
 */
type CometRegister struct {
	CometId string //comet ID
	TcpAddr string //comet对外开放tcp服务地址
	WsAddr  string //comet对外开放ws服务地址
	RpcAddr string //反连地址(comet rpc服务地址)
}

/* 鉴权
*方向:Comet->Router
 */
type AuthRequest struct {
	Id       string
	Termtype int
	Code     string
}

/* 用户socket状态通知
* 方向:Comet->Router
 */
type StateNotify struct {
	Id       string
	Termtype int
	Token    string
	CometId  string //附着Comet ID
	State    int    //1-online 0-offline
}

/*推送请求
* 方向:router->comet
 */
type MsgUpwardRequst struct {
	Id       string
	Termtype int
	Msg      string
}

/* 踢人下线
*方向:router->comet
 */
type KickRequst struct {
	Id       string
	Termtype int
	Token    string
	/* added by liang @ 2016-07-11
	0x01:重复登录，同终端类型客户端登录
	0x02:互斥登录，Android/iOS互斥
	0x03:session超时
	*/
	Reason int
}

/*推送请求
* 方向:router->comet
 */
type PushRequst struct {
	Tp         int //消息类型
	Flag       int //IOS 声音提示
	Id         string
	Termtype   int
	AppleToken string
	Msg        string
}

/* 心跳检测
* 方向：RPC客户端->RPC服务端
 */
type Ping struct {
}

//公共应答
type Response struct {
	Code int //0-成功 -1-失败
}
