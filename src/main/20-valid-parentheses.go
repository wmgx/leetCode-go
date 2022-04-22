package main

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()[]{}("))
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	var stack = make([]byte, 0)
	for _, c := range s {
		if c == ')' || c == ']' || c == '}' {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop == '(' && c == ')' {
				continue
			} else if pop == '[' && c == ']' {
				continue
			} else if pop == '{' && c == '}' {
				continue
			} else {
				return false
			}
		} else {
			stack = append(stack, byte(c))
		}
	}
	return len(stack) == 0
}
func isValid1(s string) bool {
	newStack := NewStack()
	for _, c := range s {
		if c == ')' || c == ']' || c == '}' {
			if newStack.IsEmpty() {
				return false
			}
			pop, _ := newStack.Pop()
			if pop == '(' || c == ')' {
				continue
			} else if pop == '[' || c == ']' {
				continue
			} else if pop == '{' || c == '}' {
				continue
			} else {
				return false
			}
		} else {
			newStack.Push(c)
		}
	}
	return newStack.IsEmpty()
}
