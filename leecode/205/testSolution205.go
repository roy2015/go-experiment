//205. 同构字符串
//给定两个字符串 s 和 t，判断它们是否是同构的。
//
//如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
//
//每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
//
//
//
//示例 1:
//
//输入：s = "egg", t = "add"
//输出：true
//示例 2：
//
//输入：s = "foo", t = "bar"
//输出：false
//示例 3：
//
//输入：s = "paper", t = "title"
//输出：true
//
//
//提示：
//
//可以假设 s 和 t 长度相同。

package main

import "fmt"

/**
map的使用

执行结果：通过
执行用时：4 ms, 在所有 Go 提交中击败了75.63%的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了12.66%的用户
通过测试用例：43 / 43
*/
func isIsomorphic(s string, t string) bool {
	sBytes := []byte(s)
	tBytes := []byte(t)
	frontMap := make(map[byte]byte)
	backMap := make(map[byte]byte)

	length := len(sBytes)
	for i := 0; i < length; i++ {
		sChar := sBytes[i]
		tChar := tBytes[i]
		_, existZ := frontMap[sChar]
		if !existZ {
			frontMap[sChar] = tChar
		} else {
			if frontMap[sChar] != tChar {
				return false
			}
		}
		_, existF := backMap[tChar]
		if !existF {
			backMap[tChar] = sChar
		} else {
			if backMap[tChar] != sChar {
				return false
			}
		}
	}
	return true
}

func main() {

	fmt.Println(isIsomorphic("badc", "baba")) //false
	fmt.Println(isIsomorphic("foo", "bar"))   //false
	fmt.Println(isIsomorphic("egg", "add"))   //true
	fmt.Println(isIsomorphic("foo", "bar"))   //false

}
