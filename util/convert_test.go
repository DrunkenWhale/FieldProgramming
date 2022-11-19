package util

import "testing"

func TestCNConvertToDigit(t *testing.T) {
	//t.Log(CNConvertToDigit("柒万零贰佰零叁元整"))
	//t.Log(CNConvertToDigit("零元整"))
	println(CNConvertToDigit("零元整"))
	println(CNConvertToDigit("伍仟万零伍元整"))
	println(CNConvertToDigit("壹佰零贰亿零叁佰零肆万零伍佰零陆元整"))
	println(CNConvertToDigit("壹佰亿零贰仟零叁万零肆拾元整"))
	println(CNConvertToDigit("壹佰亿零贰佰万零叁仟零肆拾元整"))
	println(CNConvertToDigit("亿伍亿伍万伍亿万仟元整"))
}
