package util

import "strconv"

type InputType = int

const (
	UNKNOWN InputType = iota
	CN
	DIGITS
)

// JudgeInputType 不满足要求的置为无法识别
func JudgeInputType(input string) InputType {
	if _, err := strconv.ParseInt(input, 10, 64); err == nil {
		// 可以被解析
		// 是数字
		return DIGITS
	}
	if len(input)%3 != 0 {
		// 不可能是汉字
		// gbk是三个字节的
		return UNKNOWN
	}
	return UNKNOWN
}
