package DigitalSwitching

import "unicode"

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

func switching(digits string) string { //组内转换
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
	return Digits
}

func Switching(digits string) string {
	for _, digit := range digits {
		if !unicode.IsDigit(digit) {
			return "输入错误"
		}
	}
	var Digits string
	if len(digits) > 8 { //亿
		part := digits[0 : len(digits)-8]
		partDigits := switching(part)
		Digits = connectDigits(Digits, partDigits)
		if partDigits != "零" {
			Digits = Digits + "亿"
		}
	}
	for len(digits) > 8 {
		digits = digits[1:len(digits)]
	}
	if len(digits) > 4 { //万
		part := digits[0 : len(digits)-4]
		partDigits := switching(part)
		Digits = connectDigits(Digits, partDigits)
		if partDigits != "零" {
			Digits = Digits + "万"
		}
	}
	for len(digits) > 4 {
		digits = digits[1:len(digits)]
	}
	partDigits := switching(digits)
	Digits = connectDigits(Digits, partDigits)
	if len(Digits) > 3 && firstDigit(Digits) == "零" {
		Digits = Digits[3:]
	}
	if firstDigit(Digits) != "零" && lastDigit(Digits) == "零" {
		Digits = Digits[:len(Digits)-3]
	}
	Digits = Digits + "元整"
	return Digits
}
