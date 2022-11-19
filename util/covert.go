package util

import (
	"strings"
)

var keyWords = []string{
	"亿",
	"万",
	"元整",
}

func CNConvertToDigit(input string) int {
	res := 0
	input = strings.ReplaceAll(input, "零", "")
	if 1 == strings.Count(input, keyWords[0]) {
		// 大于1亿
		arr := strings.Split(input, keyWords[0])
		i := PartCNToDigit(arr[0])
		input = arr[1]
		if i == -1 {
			return -1
		}
		res += 1e8 * i
	}
	if 1 == strings.Count(input, keyWords[1]) {
		// 大于1w
		arr := strings.Split(input, keyWords[1])
		i := PartCNToDigit(arr[0])
		input = arr[1]
		if i == -1 {
			return -1
		}
		res += 1e4 * i
	}
	if 1 == strings.Count(input, keyWords[2]) {
		// 大于1
		arr := strings.Split(input, keyWords[2])
		i := PartCNToDigit(arr[0])
		input = arr[1]
		if i == -1 {
			return -1
		}
		res += i
	}
	return res
}

var partKeyWords = []string{
	"仟",
	"佰",
	"拾",
}

var CNToDigitMap = map[string]int{
	"壹": 1,
	"贰": 2,
	"叁": 3,
	"肆": 4,
	"伍": 5,
	"陆": 6,
	"柒": 7,
	"捌": 8,
	"玖": 9,
}

func PartCNToDigit(part string) int {
	part = strings.ReplaceAll(part, "零", "")
	res := 0
	if 1 == strings.Count(part, partKeyWords[0]) {
		arr := strings.Split(part, partKeyWords[0])
		part = arr[1]
		s := arr[0]
		i, b := CNToDigitMap[s]
		if !b {
			// 不在映射表中
			return -1
		}
		res += i * 1000
	}
	if 1 == strings.Count(part, partKeyWords[1]) {
		arr := strings.Split(part, partKeyWords[1])
		part = arr[1]
		s := arr[0]
		i, b := CNToDigitMap[s]
		if !b {
			// 不在映射表中
			return -1
		}
		res += i * 100
	}
	if 1 == strings.Count(part, partKeyWords[2]) {
		arr := strings.Split(part, partKeyWords[2])
		part = arr[1]
		s := arr[0]
		i, b := CNToDigitMap[s]
		if !b {
			// 不在映射表中
			return -1
		}
		res += i * 10
	}
	if len(part) == 0 {
		return res
	}
	i, b := CNToDigitMap[part[len(part)-3:]]
	if !b {
		return res
	}
	res += i
	return res
}

func DigitConvertToCN(input int) string {
	return ""
}
