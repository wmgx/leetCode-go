package main

import "fmt"

/**
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x/2

	for left < right {
		mid := left + (right-left+1)/2
		value := mid * mid
		if value == x {
			return mid
		}
		// 小于的时候可能是它也可能不是它
		if value < x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
