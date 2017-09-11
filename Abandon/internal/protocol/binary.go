package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

//GetMsgType 获取指定头中消息类型
func GetMsgType(head *Header) int {
	msg := head.Line & 0x3F
	return int(msg)
}

//SetMsgType 设置指定头中消息类型
func SetMsgType(head *Header, msgType int) {
	head.Line = head.Line >> 6
	head.Line = head.Line << 6
	head.Line = head.Line | byte(msgType)
}

//SetMsgType 获取加密方式
func GetEncode(head *Header) int {
	encode := head.Line & 0xC0 >> 6
	return int(encode)
}

//SetMsgType 设置加密方式
func SetEncode(head *Header, encode int) {
	head.Line = head.Line << 2
	head.Line = head.Line >> 2
	head.Line = head.Line | byte(encode)<<6
}

//EncodeHeader 对固定头序列化
func EncodeHeader(head *Header) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, head)
	return buf.Bytes()
}

//EncodeAddHeader 对附加头进行序列化
func EncodeAddHeader(addhead *AddHeader) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, addhead)
	return buf.Bytes()
}

//EncodeBody 对消息体进行序列化
func EncodeBody(body interface{}) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, body)
	return buf.Bytes()
}

//DecodeHeader 协议固定头解码
func DecodeHeader(data []byte) (*Header, error) {
	var head Header
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &head)
	return &head, err
}

//DecodeAddHeader 协议附加头解码
func DecodeAddHeader(data []byte) (*AddHeader, error) {
	var head AddHeader
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &head)
	return &head, err
}

//DecodeBody 消息体解码
func DecodeBody(data []byte) (interface{}, error) {
	var body interface{}
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &body)
	return &body, err
}

//EncodeParamUaReg 对用户注册消息编码
func EncodeParamReg(reg *ParamReg) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, reg)
	return buf.Bytes()
}

//DecodeParamUaReg 对用户注册消息解码
func DecodeParamReg(data []byte) (*ParamReg, error) {
	var reg ParamReg
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &reg)
	return &reg, err
}

//EncodeParamUaPush 对Push参数消息解码
func EncodeParamPush(msg *ParamPush) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, msg)
	return buf.Bytes()
}

func DecodeParamPush(data []byte) (*ParamPush, error) {
	var msg ParamPush
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &msg)
	return &msg, err
}

func EncodeParamResp(msg *ParamResp) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, msg)
	return buf.Bytes()
}

func DecodeParamResp(data []byte) (*ParamResp, error) {
	var msg ParamResp
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &msg)
	return &msg, err
}

//CheckBuffer 检查二进制序列长度是否合法
func CheckBuffer(buffer []byte, min int) error {
	l := len(buffer)
	if l < min {
		return errors.New("NOT RIGHT FORMAT")
	}
	return nil
}

//buffer移字节
func ShiftBuffer(buffer []byte, length int) []byte {
	buffer = buffer[length:]
	return buffer
}

func packBinary(data interface{}) []byte {
	var buf []byte
	switch data.(type) {
	case *Register:
		msg := data.(*Register)
		b1 := bigEndian(msg.Header)
		buf = append(buf, b1...)
		b2 := bigEndian(msg.AddHeader)
		buf = append(buf, b2...)
		b3 := bigEndian(msg.ParamReg)
		buf = append(buf, b3...)
		CodecEncode(buf[HEADER_LEN:], int(msg.Len), GetEncode(&msg.Header))
	case *Push:
		msg := data.(*Push)
		b1 := bigEndian(msg.Header)
		buf = append(buf, b1...)
		b2 := bigEndian(msg.AddHeader)
		buf = append(buf, b2...)
		b3 := bigEndian(msg.ParamPush)
		buf = append(buf, b3...)
		b4 := bigEndian(msg.Msg)
		buf = append(buf, b4...)
		CodecEncode(buf[HEADER_LEN:], int(msg.Len), GetEncode(&msg.Header))
	case *Callback:
		msg := data.(*Callback)
		b1 := bigEndian(msg.Header)
		buf = append(buf, b1...)
		b2 := bigEndian(msg.AddHeader)
		buf = append(buf, b2...)

		b4 := bigEndian(msg.Msg)
		buf = append(buf, b4...)
		CodecEncode(buf[HEADER_LEN:], int(msg.Len), GetEncode(&msg.Header))
	case *ImDown:
		msg := data.(*ImDown)
		b1 := bigEndian(msg.Header)
		buf = append(buf, b1...)
		b2 := bigEndian(msg.AddHeader)
		buf = append(buf, b2...)
		b3 := bigEndian(msg.Flag)
		buf = append(buf, b3...)
		b4 := bigEndian(msg.Msg)
		buf = append(buf, b4...)
		CodecEncode(buf[HEADER_LEN:], int(msg.Len), GetEncode(&msg.Header))
	case *ImUp:
		msg := data.(*ImUp)
		b1 := bigEndian(msg.Header)
		buf = append(buf, b1...)
		b2 := bigEndian(msg.AddHeader)
		buf = append(buf, b2...)
		b4 := bigEndian(msg.Msg)
		buf = append(buf, b4...)
		CodecEncode(buf[HEADER_LEN:], int(msg.Len), GetEncode(&msg.Header))
	case *Resp:
		msg := data.(*Resp)
		b1 := bigEndian(msg.Header)
		buf = append(buf, b1...)
		b2 := bigEndian(msg.AddHeader)
		buf = append(buf, b2...)

		b3 := bigEndian(msg.Code)
		buf = append(buf, b3...)
		CodecEncode(buf[HEADER_LEN:], int(msg.Len), GetEncode(&msg.Header))
	case *Kick:
		msg := data.(*Kick)
		b1 := bigEndian(msg.Header)
		buf = append(buf, b1...)
		b2 := bigEndian(msg.AddHeader)
		buf = append(buf, b2...)

		b3 := bigEndian(msg.Reason)
		buf = append(buf, b3...)
		CodecEncode(buf[HEADER_LEN:], int(msg.Len), GetEncode(&msg.Header))
	case *Header:
		buf = bigEndian(data)
	default:
	}

	//buf = bigEndian(data)
	return buf
}

//大端格式二进制序列化
func bigEndian(data interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}
