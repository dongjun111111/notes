package protocol

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"galopush/internal/logs"
)

func packJson(v interface{}) []byte {
	var buf []byte
	var err error
	var body map[string]interface{}
	body = make(map[string]interface{})
	var data map[string]interface{}
	data = make(map[string]interface{})
	switch v.(type) {
	case *Push:
		msg := v.(*Push)
		body["cmd"] = GetMsgType(&msg.Header)
		body["tid"] = msg.Tid
		data["msg"] = string(msg.Msg)
		body["data"] = data
		buf, err = json.Marshal(&body)
		if err != nil {
			logs.Logger.Error(err)
		}
	case *Callback:
		msg := v.(*Callback)
		body["cmd"] = GetMsgType(&msg.Header)
		body["tid"] = msg.Tid
		data["msg"] = string(msg.Msg)
		body["data"] = data
		buf, err = json.Marshal(&body)
		if err != nil {
			logs.Logger.Error(err)
		}
	case *ImDown:
		msg := v.(*ImDown)
		body["cmd"] = GetMsgType(&msg.Header)
		body["tid"] = msg.Tid
		data["msg"] = string(msg.Msg)
		body["data"] = data
		buf, err = json.Marshal(&body)
		if err != nil {
			logs.Logger.Error(err)
		}
	case *Resp:
		msg := v.(*Resp)
		body["cmd"] = GetMsgType(&msg.Header)
		body["tid"] = msg.Tid
		data["code"] = int(msg.Code)
		body["data"] = data
		buf, err = json.Marshal(&body)
		if err != nil {
			logs.Logger.Error(err)
		}
	case *Kick:
		msg := v.(*Kick)
		body["cmd"] = GetMsgType(&msg.Header)
		body["tid"] = msg.Tid
		data["reson"] = int(msg.Reason)
		body["data"] = data
		buf, err = json.Marshal(&body)
		if err != nil {
			logs.Logger.Error(err)
		}
	}
	logs.Logger.Debug("packJson=", string(buf[:]))
	return []byte(Encode(buf))
}

func marshal(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	return b, err
}

func Encode(buff []byte) string {
	CodecEncode(buff, len(buff), ENCODE_LOOP_XOR)
	return base64.StdEncoding.EncodeToString(buff)
}

func Decode(buff []byte) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(string(buff))
	if err != nil {
		logs.Logger.Error("Decode ws on decode base64 error ", err)
		return b, err
	}
	CodecDecode(b, len(b), ENCODE_LOOP_XOR)
	return b, nil
}

func UnPackJson(buffer []byte) (interface{}, error) {
	b, err := Decode(buffer)
	if err != nil {
		return nil, err
	}
	var body map[string]interface{}
	if err := json.Unmarshal(b, &body); err != nil {
		logs.Logger.Error(err, " msg=", string(b))
		return nil, err
	}
	cmd := int(body["cmd"].(float64))
	tid := int(body["tid"].(float64))
	data := body["data"].(map[string]interface{})
	switch cmd {
	case MSGTYPE_REGISTER:
		{
			var msg Register
			SetMsgType(&msg.Header, cmd)
			SetEncode(&msg.Header, 0)
			msg.Tid = uint32(tid)
			msg.Len = 66

			//消息体
			msg.Version = byte(data["version"].(float64))
			msg.TerminalType = byte(data["termType"].(float64))
			idBuff := []byte(data["id"].(string))
			for i := 0; i < 32 && i < len(idBuff); i++ {
				msg.Id[i] = idBuff[i]
			}
			tokenBuff := []byte(data["token"].(string))
			for i := 0; i < 32 && i < len(tokenBuff); i++ {
				msg.Token[i] = tokenBuff[i]
			}
			//p.Cache(conn, &msg)
			return &msg, nil
		}
	case MSGTYPE_HEARTBEAT:
		{
			var msg Header
			//固定头
			SetMsgType(&msg, cmd)
			SetEncode(&msg, 0)
			msg.Tid = uint32(tid)
			return &msg, nil
		}

	case MSGTYPE_PUSHRESP, MSGTYPE_CBRESP, MSGTYPE_MSGRESP, MSGTYPE_KICKRESP:
		{
			var msg Resp
			//固定头
			SetMsgType(&msg.Header, cmd)
			SetEncode(&msg.Header, 0)
			msg.Tid = uint32(tid)

			//可变头
			msg.Len = 1
			msg.Code = byte(data["code"].(float64))
			return &msg, nil
		}
	case MSGTYPE_MESSAGE:
		var msg ImUp
		//固定头
		SetMsgType(&msg.Header, cmd)
		SetEncode(&msg.Header, 0)
		msg.Tid = uint32(tid)

		content := data["msg"].(string)
		contentBuff := []byte(content)

		msg.Len = uint32(len(contentBuff))

		msg.Msg = append(msg.Msg, contentBuff...)

		return &msg, nil
	}

	return nil, errors.New("UnKnow cmd type")
}
