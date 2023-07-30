package idBuilder

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var tenToAny = []byte("9I5KLH4QJ7YXDTMNPWFA8ECGUVR3S2B6")

// 10进制转32进制
func from10To32(num int) string {
	newNumStr := ""
	var remainder int
	var remainderString string
	for num != 0 {
		remainder = num % 32
		remainderString = string(tenToAny[remainder])
		newNumStr = remainderString + newNumStr
		num = num / 32
	}
	return newNumStr
}

// From32To10 32进制转10进制 （根据邀请码生成用户的uuid）
func From32To10(num string) int {
	dict := string(tenToAny) + "0"
	var newNum float64
	newNum = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(strings.Index(dict, value))
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(32), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(newNum)
}

// Id2Code 生成邀请码 num为自增的用户id
func Id2Code(num int) string {
	code := from10To32(num)
	codeLen := len(code)
	if (6 - codeLen) > 0 {
		code = code + "0"
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 5 - codeLen; 0 < i; i-- {
			code = code + string(tenToAny[r.Intn(32)])
		}
	}
	return code
}
