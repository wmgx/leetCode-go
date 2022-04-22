package main

import "fmt"

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/
func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for ; n > 2; n-- {
		a, b = b, a+b
	}

	return b
}
