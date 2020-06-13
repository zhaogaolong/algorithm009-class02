package week04

import "sort"

type solution struct {
	max    int
	amount int
	coins  []int
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	s := solution{
		max:    -1,
		amount: amount,
		coins:  coins,
	}
	s.dfs(0, 0, 0)
	return s.max
}

func (this *solution) dfs(curr, index, counter int) {
	if curr == this.amount {
		if counter < this.max || this.max < 0 {
			this.max = counter
		}
		return
	}

	for i := index; i <= 0; i-- {
		if curr+this.coins[i] > this.amount {
			continue
		}

		// max > 0 已经说明有结果了，接下来判断是不是有更好的结果，如果有就继续，没有就结束当前的迭代
		//
		// (this.amount - curr) / this.coins[i] 是剩余的钱还能使用多少个当前面值的硬币
		// this.max - counter 当前的结果（max）减去已经继续使用了的硬币数量（在本次递归中）
		// 当前的面值使用的数量已经大于之前的解，就是没有之前的结果（max - counter） 好，直接就终止循环递归
		if this.max > 0 && (this.amount-curr)/this.coins[i] > this.max-counter {
			break
		}
		this.dfs(curr+this.coins[i], index, counter+1)
	}
}
