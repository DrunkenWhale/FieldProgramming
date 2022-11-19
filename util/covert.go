package util

import "strings"

var keyWords = []string{
	"亿",
	"万",
	"元整",
}

func CNConvertToDigit(input string) int {
	if 1 == strings.Count(input, keyWords[0]) {
		// 大于1亿

	}
	return 0
}

func DigitConvertToCN(input int) string {
	return ""
}
