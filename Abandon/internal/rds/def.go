package rds

type Sessions struct {
	Id      string     `json:"id"`
	CometId string     `json:"comet"`
	Sess    []*Session `json:"sess"`
}

type Session struct {
	Id       string `json:"id"`
	Plat     int    `json:"plat"`
	Online   bool   `json:"online"`
	Login    bool   `json:"login"`
	AuthCode string `json:"authcode"`
	IOSToken string `json:"iosToken"`
	OffMsg   Body   `json:"offmsg"`
}

type Body struct {
	Push Push    `json:"push"`
	CB   []*Item `json:"callback"`
	Im   []*Item `json:"im"`
}

type Item struct {
	WebOnline bool   `json:"webOnline"`
	Msg       []byte `json:"msg"`
}

type Push struct {
	OffCnt uint16 `json:"count"`
	Msg    []byte `json:"msg"`
}
