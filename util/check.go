package util

type InputType = int

const (
	UNKNOWN InputType = iota
	CN
	DIGITS
)

// JudgeInputType 不满足要求的置为无法识别
func JudgeInputType(input string) InputType {
	return UNKNOWN
}
