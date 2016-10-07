package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

var sem = make(chan int)
var yidong []string = []string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188"}
var liantong []string = []string{"130", "131", "132", "145", "155", "156", "176", "185", "186"}
var dianxin []string = []string{"133", "153", "177", "180", "181", "189", "173"}
var qita []string = []string{"170"} //1700,1705,1709

func RandInt64(min, max int64) int64 {
JASON:
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandInt64(min, max)
		goto JASON
	}
	return i.Int64()
}

func main() {
	fmt.Println(RandInt64(1000, 9999))
	var allnumberprefix []string
	allnumberprefix = append(allnumberprefix, yidong...)
	allnumberprefix = append(allnumberprefix, liantong...)
	allnumberprefix = append(allnumberprefix, dianxin...)
	allnumberprefix = append(allnumberprefix, qita...)
	fmt.Println(allnumberprefix[int(RandInt64(0, int64(len(allnumberprefix)-1)))] + strconv.Itoa(int(RandInt64(10000000, 99999999))))
}
