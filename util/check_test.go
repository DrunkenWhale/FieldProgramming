package util

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestJudgeInputType(t *testing.T) {
	assert.Equal(t, JudgeInputType("1145141919810"), DIGITS)
	assert.Equal(t, JudgeInputType("差不多的了"), UNKNOWN)
}
