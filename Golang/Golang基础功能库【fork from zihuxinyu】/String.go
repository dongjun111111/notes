package Library

import (
	"regexp"
	"strings"
	"strconv"
)

var (
	lowerRe, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	brRe, _    = regexp.Compile("<br.*?>")
	tagRe, _   = regexp.Compile("<.*?>")
)

func StripTags(html string) string {
	html = lowerRe.ReplaceAllStringFunc(html, strings.ToLower)
	html = strings.Replace(html, "\n", " ", -1)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "&nbsp;", " ", -1)
	html = brRe.ReplaceAllString(html, "")
	html = tagRe.ReplaceAllString(html, "")
	return html
}

// Simplify HTML text by removing tags
func RemoveFormatting(html string) string {
	return StripTags(html)
}

//Html过滤
func Html2str(html string) string {
	src := string(html)
	//替换HTML的空白字符为空格
	re := regexp.MustCompile(`\s`) //ns*r
	src = re.ReplaceAllString(src, " ")
	//将HTML标签全转换成小写
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

// 以byte来截取
func SubstringByte(str string, start int) string {
	return substr(str, start, len(str)-start, false)
}
func Substring(str string, start int) string {
	return substr(str, start, len(str)-start, true)
}
func Substr(str string, start, length int) string {
	return substr(str, start, length, true)
}
func substr(str string, start, length int, isRune bool) string {
	rs := []rune(str)
	rs2 := []byte(str)
	rl := len(rs)
	if !isRune {
		rl = len(rs2)
	}
	end := 0
	if start < 0 {
		start = rl-1+start
	}
	end = start+length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	if isRune {
		return string(rs[start:end])
	}
	return string(rs2[start:end])
}

//列表是否包含给定项
func ListContains(list []interface{}, key interface{}) (finded bool) {
	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}

//字符串数组中是否包含给定项
func StringsContains(list []string, key string) (finded bool) {
	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}


func StringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

//字符串转长整型
func Str2int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

//字符串转整形
func Str2int(s string) (int, error) {
	return strconv.Atoi(s)
}

//整形转字符串
func Int2str(i int) string {
	return strconv.Itoa(i)
}

// convert like this: "HelloWorld" to "hello_world"
func SnakeCasedName(name string) string {
	newstr := make([]rune, 0)
	firstTime := true
	for _, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if firstTime == true {
				firstTime = false
			} else {
				newstr = append(newstr, '_')
			}
			chr -= ('A'-'a')
		}
		newstr = append(newstr, chr)
	}
	return string(newstr)
}

//将数组转换为字符串，英文逗号分隔
func Array2String(array []string) string {
	item := ""
	for _, v := range array {
		item+=v+","
	}
	item = item[:len(item)-1]
	return item
}

