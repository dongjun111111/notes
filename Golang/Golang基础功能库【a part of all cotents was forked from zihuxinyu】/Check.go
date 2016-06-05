package Library

import "regexp"

// 是否是email
func IsEmail(email string) bool {
	if email == "" {
		return false;
	}
	ok, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[0-9a-zA-Z]{2,3}$`, email)
	return ok
}

// 是否只包含数字, 字母 -, _
func IsUsername(username string) bool {
	if username == "" {
		return false;
	}
	ok, _ := regexp.MatchString(`[^0-9a-zA-Z_\-]`, username)
	return !ok
}

///检测是否为map[string]string类型
func IsMap_String_String(obj interface{}) bool {
	switch obj.(type){

	case map[string]string:

		return true

	default:

		return false

	}

}
