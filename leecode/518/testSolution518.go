package main

import (
	"fmt"
)

//518. 零钱兑换 II
//给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//
//请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
//假设每一种面额的硬币有无限个。
//
//题目数据保证结果符合 32 位带符号整数。
//
//
//
//示例 1：
//
//输入：amount = 5, coins = [1, 2, 5]
//输出：4
//解释：有四种方式可以凑成总金额：
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
//示例 2：
//
//输入：amount = 3, coins = [2]
//输出：0
//解释：只用面额 2 的硬币不能凑成总金额 3 。
//示例 3：
//
//输入：amount = 10, coins = [10]
//输出：1
//
//
//提示：
//
//1 <= coins.length <= 300
//1 <= coins[i] <= 5000
//coins 中的所有值 互不相同
//0 <= amount <= 5000

type solution struct {
	amount int
	coins  []int
}

func (so *solution) change1(amount int) int {
	if amount == 0 {
		return 1
	} else if amount < 0 {
		return 0
	}
	k := 0

	for _, i := range so.coins {
		k += so.change1(amount - i)
	}
	return k
}

/**
虽然未通过单学会了面向对象
*/
func change1(amount int, coins []int) int {
	so := solution{amount: amount, coins: coins}
	return so.change1(amount)
}

/**
执行结果：通过
显示详情
添加备注

执行用时：16 ms, 在所有 Go 提交中击败了6.14%的用户
内存消耗：12.6 MB, 在所有 Go 提交中击败了5.12%的用户通过测试用例：28 / 28
*/
func change(amount int, coins []int) int {
	var k = len(coins)
	dp := make([][]int, k)
	col := amount + 1
	row := k

	//初始化数组,分配空间
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	//初始化第一行
	for i := 0; i < col; i++ {
		if (i % coins[0]) == 0 {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}

	for i := 1; i < row; i++ {
		var coin = coins[i]
		dp[i][0] = 1
		for j := 1; j < col; j++ {
			var r2 = 0
			var r1 = dp[i-1][j]
			for k := j - coin; k >= 0; k -= coin {
				r2 += dp[i-1][k]
			}
			dp[i][j] = r1 + r2
		}
	}
	return dp[k-1][amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5})) //4
	fmt.Println(change(3, []int{2}))       //0
	fmt.Println(change(10, []int{10}))     //1
}
