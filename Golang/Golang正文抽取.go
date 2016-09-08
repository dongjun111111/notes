package main

import (
	. "fmt"
	"io/ioutil"
	"net/http"
	_ "os"
	"regexp"
	"strings"
)

const blkSize int = 3

var (
	lines   []string
	blksLen []int
	isGB    bool
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Get(url string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			Println(err)
		}
	}()
	resp, err := http.Get(url)
	check(err)
	//Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		panic("FUCK")
	}
	return ioutil.ReadAll(resp.Body)
}

func strip(src string) string {
	src = strings.ToLower(src)
	re, _ := regexp.Compile(`<!doctype.*?>`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`<!--.*?-->`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`<.*?>`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`&.{1,5};|&#.{1,5};`)
	src = re.ReplaceAllString(src, "")

	src = strings.Replace(src, "\r\n", "\n", -1)
	src = strings.Replace(src, "\r", "\n", -1)
	return src
}

func parse(src string) {
	array := strings.Split(src, "\n")
	for _, a := range array {
		a = strings.Replace(a, " ", "", -1)
		lines = append(lines, a)
	}

	blen := 0
	for i := 0; i < blkSize; i++ {
		blen += len(lines[i])
	}
	blksLen = append(blksLen, blen)
	for i := 1; i < len(lines)-blkSize; i++ {
		blen = blksLen[i-1] + len(lines[i-1+blkSize]) - len(lines[i-1])
		blksLen = append(blksLen, blen)
	}
}

func Do(url string) string {
	body, err := Get(url)
	check(err)
	parse(strip(string(body)))
	i, max := 0, 0
	blkNum := len(blksLen)
	plainText := ""
	for i < blkNum {
		for i < blkNum && blksLen[i] == 0 {
			i++
		}
		if i > blkNum {
			break
		}
		curTextLen, portion := 0, ""
		for i < blkNum && blksLen[i] > 0 {
			if lines[i] != "" {
				portion += lines[i] + "<br />"
				curTextLen += len(lines[i])
			}
			i++
		}
		if curTextLen > max {
			plainText = portion
			max = curTextLen
		}
	}
	return plainText
}

func main() {
	plain := Do("http://localhost:2333/gameserver-gate")
	Println(plain)
}
