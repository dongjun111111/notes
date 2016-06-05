package Library

import (
	"strings"
	"github.com/astaxie/beego/config"
)

var jsons config.ConfigContainer


func Tr(key string, args ...string) string {
	//语言
	local := "zh"
	if len(args) > 0 {
		local = args[0]
	}
	//按语言读取
	lang, _ := jsons.DIY(local)
	m := lang.(map[string]interface{})
	for k, v := range m {
		if k == key {
			return v.(string)
		}
	}
	return ""
}

/*
根据客户端语言环境，取得相关语言字符串
key
lang：取自客户端的Accept-language
*/
func Lang(key, lang string) string {
	return Tr(key, Local(lang))
}

/*
根据客户端语言环境，确定客户端语言
lang：取自客户端的Accept-language
*/
func Local(lang string) string {
	switch {
	case strings.Contains(lang, "zh"):
		return "zh"
	case strings.Contains(lang, "en"):
		return "en"
	default:
		return "zh"
	}
}
