//package main
//
//import "fmt"
//
//func main()  {
//	//函数的基本操作：声明与相关功能验证；多个返回值
//
//	//可变参数
//
//	//固定参数与可变参数
//
//	//defer语句
//
//	//全局变量与局部变量
//
//	//函数作为变量
//
//	//函数作为参数
//
//	//匿名函数：变量声明；直接使用
//
//	//闭包：=匿名函数+环境；常见闭包
//
//	//panic：基础定义；与defer、recover（）相关的操作；常见操作分析
//}

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
)

const (
	win            = 100 // The winning score in a game of Pig
	gamesPerSeries = 10  // The number of games per series to simulate
)

// A score includes scores accumulated in previous turns for each player,
// as well as the points scored by the current player in this turn.
type score struct {
	player, opponent, thisTurn int
}

/*在Go中，函数可以像其它值那样到处传递。一个函数的类型签名描述了其实参与返回值的类型。
action 是一个函数，它接受一个 score 并返回产生的 score 与当前一轮是否结束。*/
type action func(current score) (result score, turnIsOver bool)

/*Go的函数可以返回多个值
函数 roll 和 stay 都返回一对值。它们也匹配 action 的类型签名。这些 action 类型的函数定义了Pig的规则。*/
// roll returns the (result, turnIsOver) outcome of simulating a die roll.
// If the roll value is 1, then thisTurn score is abandoned, and the players'
// roles swap.  Otherwise, the roll value is added to thisTurn.
func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1 // A random int in [1, 6]
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.opponent, outcome + s.thisTurn}, false
}

// stay returns the (result, turnIsOver) outcome of staying.
// thisTurn score is added to the player's score, and the players' roles swap.
func stay(s score) (score, bool) {
	return score{s.opponent, s.player + s.thisTurn, 0}, true
}

/*一个函数能够将其它函数作为实参和返回值。*/
// A strategy chooses an action for any given score.
type strategy func(score) action

/*Go中可以声明匿名函数。函数字面是闭包的： 它们继承了声明它们所在函数的作用域。*/
// stayAtK returns a strategy that rolls until thisTurn is at least k, then stays.
func stayAtK(k int) strategy {
	return func(s score) action {
		if s.thisTurn >= k {
			return stay
		}
		return roll
	}
}

// play simulates a Pig game and returns the winner (0 or 1).
func play(strategy0, strategy1 strategy) int {
	strategies := []strategy{strategy0, strategy1}
	var s score
	var turnIsOver bool
	// 随机决定谁先玩
	currentPlayer := rand.Intn(2) // Randomly decide who plays first
	for s.player+s.thisTurn < win {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			currentPlayer = (currentPlayer + 1) % 2
		}
	}
	return currentPlayer
}

// roundRobin simulates a series of games between every pair of strategies.
func roundRobin(strategies []strategy) ([]int, int) {
	wins := make([]int, len(strategies))
	for i := 0; i < len(strategies); i++ {
		for j := i + 1; j < len(strategies); j++ {
			for k := 0; k < gamesPerSeries; k++ {
				winner := play(strategies[i], strategies[j])
				if winner == 0 {
					wins[i]++
				} else {
					wins[j]++
				}
			}
		}
	}
	// 不能自己一个人玩
	gamesPerStrategy := gamesPerSeries * (len(strategies) - 1) // no self play
	return wins, gamesPerStrategy
}

/*像 ratioString 这样的变参函数接受数量可变的实参。
这些参数实参在该函数中可作为一个切片使用。*/
// ratioString takes a list of integer values and returns a string that lists
// each value and its percentage of the sum of all values.
// e.g., ratios(1, 2, 3) = "1/6 (16.7%), 2/6 (33.3%), 3/6 (50.0%)"
func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ", "
		}
		pct := 100 * float64(val) / float64(total)
		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	strategies := make([]strategy, win)
	for k := range strategies {
		strategies[k] = stayAtK(k + 1)
	}
	wins, games := roundRobin(strategies)

	for k := range strategies {
		fmt.Printf("Wins, losses staying at k =% 4d: %s\n",
			k+1, ratioString(wins[k], games-wins[k]))
	}
}
