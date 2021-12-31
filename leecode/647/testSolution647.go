package main

import "fmt"

//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
//回文字符串 是正着读和倒过来读一样的字符串。
//
//子字符串 是字符串中的由连续字符组成的一个序列。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
//示例 1：
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//示例 2：
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//提示：
//
//1 <= s.length <= 1000
//s 由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindromic-substrings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//执行结果：通过
//执行用时：56 ms, 在所有 Go 提交中击败了10.02%的用户
//内存消耗：2 MB, 在所有 Go 提交中击败了51.86%的用户
//通过测试用例：130 / 130
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	bytes := []byte(s)
	len := len(bytes)
	result := 0
	for i := 0; i < len; i++ {
		for j := i + 1; j <= len; j++ {
			flag := isHuiWen(bytes[i:j])
			if flag {
				result++
			}

		}
	}
	return result
}

//判断是否回文
func isHuiWen(chars []byte) bool {
	if len(chars) == 1 {
		return true
	}
	j := len(chars) - 1
	for i := 0; i < j; i++ {
		if chars[i] != chars[j] {
			return false
		}
		j--
	}
	return true
}

func main() {
	fmt.Println(countSubstrings("aba")) //4
	fmt.Println(countSubstrings("abc")) //3
	fmt.Println(countSubstrings("aaa")) //6

}
