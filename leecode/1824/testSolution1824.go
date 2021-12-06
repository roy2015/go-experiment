package main

import (
	"fmt"
	"math"
)

/*1824. 最少侧跳次数
给你一个长度为 n 的 3 跑道道路 ，它总共包含 n + 1 个 点 ，编号为 0 到 n 。一只青蛙从 0 号点第二条跑道 出发 ，它想要跳到点 n 处。然而道路上可能有一些障碍。

给你一个长度为 n + 1 的数组 obstacles ，其中 obstacles[i] （取值范围从 0 到 3）表示在点 i 处的 obstacles[i] 跑道上有一个障碍。如果 obstacles[i] == 0 ，那么点 i 处没有障碍。任何一个点的三条跑道中 最多有一个 障碍。

比方说，如果 obstacles[2] == 1 ，那么说明在点 2 处跑道 1 有障碍。
这只青蛙从点 i 跳到点 i + 1 且跑道不变的前提是点 i + 1 的同一跑道上没有障碍。为了躲避障碍，这只青蛙也可以在 同一个 点处 侧跳 到 另外一条 跑道（这两条跑道可以不相邻），但前提是跳过去的跑道该点处没有障碍。

比方说，这只青蛙可以从点 3 处的跑道 3 跳到点 3 处的跑道 1 。
这只青蛙从点 0 处跑道 2 出发，并想到达点 n 处的 任一跑道 ，请你返回 最少侧跳次数 。

注意：点 0 处和点 n 处的任一跑道都不会有障碍。



示例 1：


输入：obstacles = [0,1,2,3,0]
输出：2
解释：最优方案如上图箭头所示。总共有 2 次侧跳（红色箭头）。
注意，这只青蛙只有当侧跳时才可以跳过障碍（如上图点 2 处所示）。
示例 2：


输入：obstacles = [0,1,1,3,3,0]
输出：0
解释：跑道 2 没有任何障碍，所以不需要任何侧跳。
示例 3：


输入：obstacles = [0,2,1,0,3,0]
输出：2
解释：最优方案如上图所示。总共有 2 次侧跳。


提示：

obstacles.length == n + 1
1 <= n <= 5 * 10000
0 <= obstacles[i] <= 3
obstacles[0] == obstacles[n] == 0*/

func minSideJumps(obstacles []int) int {
	var c1 = 0
	var c2 = 0
	var c3 = 0
	for idx, obsVal := range obstacles {
		if idx == 0 {
			continue
		}
		if idx == 1 {
			switch obsVal {
			case 0:
				c1 = 1
				c3 = 1
				c2 = 0
				break
			case 1:
				c1 = math.MaxInt32
				c2 = 0
				c3 = 1
				break
			case 2:
				c1 = 1
				c2 = math.MaxInt32
				c3 = 1
				break
			case 3:
				c1 = 1
				c2 = 0
				c3 = math.MaxInt32
				break
			default:
				break
			}
			continue
		}
		switch obsVal {
		case 0:
			if c1 == math.MaxInt32 {
				var caC1 = int(math.Min(float64(c2+1), float64(c3+1)))
				c1 = caC1
			} else if c2 == math.MaxInt32 {
				var caC2 = int(math.Min(float64(c1+1), float64(c3+1)))
				c2 = caC2
			} else if c3 == math.MaxInt32 {
				var caC3 = int(math.Min(float64(c1+1), float64(c2+1)))
				c3 = caC3
			} else {
			}
			break
		case 1:
			if c1 == math.MaxInt32 {
				var caC2 = int(math.Min(float64(c2), float64(c3+1)))
				var caC3 = int(math.Min(float64(c2+1), float64(c3)))
				c2 = caC2
				c3 = caC3
			} else if c2 == math.MaxInt32 {
				var caC2 = int(math.Min(float64(c3+1), float64(c1+2)))
				var caC3 = int(math.Min(float64(c3), float64(c1+1)))
				c2 = caC2
				c3 = caC3
			} else if c3 == math.MaxInt32 {
				var caC2 = int(math.Min(float64(c2), float64(c1+1)))
				var caC3 = int(math.Min(float64(c2+1), float64(c1+2)))
				c2 = caC2
				c3 = caC3
			} else {
			}
			c1 = math.MaxInt32
			break
		case 2:
			if c1 == math.MaxInt32 {
				var caC1 = int(math.Min(float64(c2+2), float64(c3+1)))
				var caC3 = int(math.Min(float64(c2+1), float64(c3)))
				c1 = caC1
				c3 = caC3
			} else if c2 == math.MaxInt32 {
				var caC1 = int(math.Min(float64(c1), float64(c3+1)))
				var caC3 = int(math.Min(float64(c1+1), float64(c3)))
				c1 = caC1
				c3 = caC3
			} else if c3 == math.MaxInt32 {
				var caC1 = int(math.Min(float64(c1), float64(c2+1)))
				var caC3 = int(math.Min(float64(c1+1), float64(c2+2)))
				c1 = caC1
				c3 = caC3
			} else {
			}
			c2 = math.MaxInt32
			break
		case 3:
			if c1 == math.MaxInt32 {
				var caC1 = int(math.Min(float64(c2+1), float64(c3+2)))
				var caC2 = int(math.Min(float64(c2), float64(c3+1)))
				c1 = caC1
				c2 = caC2
			} else if c2 == math.MaxInt32 {
				var caC1 = int(math.Min(float64(c1), float64(c3+1)))
				var caC2 = int(math.Min(float64(c1+1), float64(c3+2)))
				c1 = caC1
				c2 = caC2
			} else if c3 == math.MaxInt32 {
				var caC1 = int(math.Min(float64(c1), float64(c2+1)))
				var caC2 = int(math.Min(float64(c1+1), float64(c2)))
				c1 = caC1
				c2 = caC2
			} else {
			}
			c3 = math.MaxInt32
			break
		default:
			break
		}
	}
	var c1c2 = int(math.Min(float64(c1), float64(c2)))
	var c1c2c3 = int(math.Min(float64(c1c2), float64(c3)))
	return c1c2c3
}

func main() {
	fmt.Println("最少横跳次数", minSideJumps([]int{0, 1, 1, 3, 3, 0})) //0
	fmt.Println("最少横跳次数", minSideJumps([]int{0, 2, 1, 0, 3, 0})) //2
	fmt.Println("最少横跳次数", minSideJumps([]int{0, 1, 2, 3, 0}))    //2

}
