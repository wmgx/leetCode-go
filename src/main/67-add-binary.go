package main

import (
	"fmt"
	"strconv"
)

/**
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0
*/

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	var ai = len(a) - 1
	var bi = len(b) - 1
	var result string
	var j uint8 = 0
	for ai >= 0 || bi >= 0 {
		if ai >= 0 {
			j += a[ai] - '0'
			ai--
		}
		if bi >= 0 {
			j += b[bi] - '0'
			bi--
		}
		result = strconv.Itoa(int(j%2)) + result
		j = j / 2
	}
	if j > 0 {
		result = strconv.Itoa(int(j)) + result
	}

	return result

}
