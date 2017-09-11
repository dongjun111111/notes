//Package protocol define the transport protocol
//定义了UA和gpush之间，以及gpush与业务层之间通信协议
//同时实现了协议编解码，加解密；协议格式详见协议文档.
package protocol

const (
	//协议版本
	CURRENT_VERSION = 0x01
	//兼容的协议版本
	COMPATIBLE_VERSION = CURRENT_VERSION
)

//协议类型枚举
const (
	PROTOCOL_TYPE_DEFAULT = 0
	PROTOCOL_TYPE_BINARY  = 1
	PROTOCOL_TYPE_JSON    = 2
)

//消息类型枚举
const (
	MSGTYPE_DEFAULT   = 0
	MSGTYPE_REGISTER  = 1
	MSGTYPE_REGRESP   = 2
	MSGTYPE_HEARTBEAT = 3
	MSGTYPE_HBRESP    = 4
	MSGTYPE_PUSH      = 5
	MSGTYPE_PUSHRESP  = 6
	MSGTYPE_CALLBACK  = 7
	MSGTYPE_CBRESP    = 8
	MSGTYPE_MESSAGE   = 9
	MSGTYPE_MSGRESP   = 10
	MSGTYPE_KICK      = 11
	MSGTYPE_KICKRESP  = 12
	MSGTYPE_MAX       = 64
)

//终端类型枚举
const (
	PLAT_DEFAULT  = 0
	PLAT_ANDROID  = 1
	PLAT_IOS      = 2
	PLAT_WINPHONE = 4
	PLAT_WEB      = 8
	PLAT_PC       = 16
	PLAT_ALL      = 31
)

//消息体加密方式枚举
const (
	ENCODE_DEFAULT  = 0
	ENCODE_BIT_NOT  = 1 //对每一个字节进行按位取反
	ENCODE_BYTE_RVS = 2 //对二进制数据进行字节交互（从首字节开始两两换位，若为单数则末字节不变）
	ENCODE_LOOP_XOR = 3 //环形异或
)

//协议头长度
const (
	FIX_HEADER_LEN = 5                               //固定头
	ADD_HEADER_LEN = 4                               //可变头
	HEADER_LEN     = FIX_HEADER_LEN + ADD_HEADER_LEN //全头
)

const (
	KICK_REASON_REPEAT  = 1
	KICK_REASON_MUTEX   = 2
	KICK_REASON_TIMEOUT = 3
)

//fix head part
type Header struct {
	Line byte
	Tid  uint32
}

//var head part
type AddHeader struct {
	Len uint32
}

//body part

// 0:public response
type ParamResp struct {
	Code byte
}

// 1:about ua
type ParamReg struct {
	Version      byte
	TerminalType byte
	Id           [32]byte
	Token        [32]byte
}

type ParamPush struct {
	Offline uint16
	Flag    uint8
}

// 0:about publice response
type Resp struct {
	Header
	AddHeader
	ParamResp
}

// 1:about ua
type Register struct {
	Header
	AddHeader
	ParamReg
}

type Push struct {
	Header
	AddHeader
	ParamPush
	Msg []byte
}

type Kick struct {
	Header
	AddHeader
	Reason uint8
}

type Callback struct {
	Header
	AddHeader
	Msg []byte
}

//即时消息上行
type ImUp struct {
	Header
	AddHeader
	Msg []byte
}

//即时消息下行
type ImDown struct {
	Header
	AddHeader
	Flag uint8
	Msg  []byte
}
