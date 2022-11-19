package util

import (
	"strings"
	"unicode"
)

var keyWords = []string{
	"亿",
	"万",
	"元整",
}

func CNConvertToDigit(input string) int {
	if input == "零元整" {
		return -2
	}
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

func DigitConvertToCN(input string) string {
	return switching(input)
}

var capitalFigures = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖", "拾"}
var digitalRight = []string{"仟", "佰", "拾", ""}

func lastDigit(Digits string) string {
	if len(Digits) == 0 {
		return ""
	}
	return Digits[len(Digits)-3:]
}

func firstDigit(Digits string) string {
	if len(Digits) == 0 {
		return ""
	}
	return Digits[:3]
}

func connectDigits(last, next string) string {
	if lastDigit(last) == "零" && firstDigit(next) == "零" {
		next = next[3:]
	}
	return last + next
}

func _switching(digits string) string { //组内转换
	var Digits string
	for len(digits) < 4 {
		digits = "0" + digits
	}
	for i := 0; i < 4; i++ {
		if lastDigit(Digits) != "零" || capitalFigures[digits[i]-'0'] != "零" {
			Digits = Digits + capitalFigures[digits[i]-'0']
		}
		if digits[i] != '0' {
			Digits = Digits + digitalRight[i]
		}
	}
	if lastDigit(Digits) == "零" {
		Digits = Digits[:len(Digits)-3]
	}
	return Digits
}

func switching(digits string) string {
	for _, digit := range digits {
		if !unicode.IsDigit(digit) {
			return "输入错误"
		}
	}
	var Digits string
	if len(digits) > 8 { //亿
		part := digits[0 : len(digits)-8]
		partDigits := _switching(part)
		Digits = connectDigits(Digits, partDigits)
		if len(partDigits) != 0 {
			Digits = Digits + "亿"
		}
	}
	for len(digits) > 8 {
		digits = digits[1:len(digits)]
	}
	if len(digits) > 4 { //万
		part := digits[0 : len(digits)-4]
		partDigits := _switching(part)
		Digits = connectDigits(Digits, partDigits)
		if len(partDigits) != 0 {
			Digits = Digits + "万"
		}
	}
	for len(digits) > 4 {
		digits = digits[1:len(digits)]
	}
	partDigits := _switching(digits)
	Digits = connectDigits(Digits, partDigits)
	if len(Digits) > 3 && firstDigit(Digits) == "零" {
		Digits = Digits[3:]
	}
	Digits = Digits + "元整"
	return Digits
}
