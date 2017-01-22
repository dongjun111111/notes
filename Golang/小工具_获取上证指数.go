package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

//获取小数四舍五入后的整数
func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

//上证指数
func GetSHIndex() string {
	resp, _ := http.Get("http://hq.sinajs.cn/list=s_sh000001")
	defer resp.Body.Close()
	var str string
	if bodyByte, err := ioutil.ReadAll(resp.Body); err == nil {
		str = string(bodyByte) //var hq_str_s_sh000001="上证指数,3453.825,19.244,0.56,905288,11706185";
	}
	strs := strings.Split(str, ",")
	set, _ := strconv.ParseFloat(strs[1], 64)
	strs2 := Round(set, 2)
	i2 := strconv.FormatFloat(strs2, 'f', 2, 64)
	id2 := i2[len(i2)-2:]
	return id2
}
func main() {
	fmt.Println("上证指数:", GetSHIndex())
}
