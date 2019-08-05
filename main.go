package main

import (
	"fmt"
	"math/rand"
	"time"
)

// balls
type Balls struct {
}

// 建立球數量
func (ball *Balls) CreateNew(maxAmountOfBalls int) []int {
	numbers := []int{}
	for i := 1; i <= maxAmountOfBalls; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

// 洗球
func (ball *Balls) Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix())) //根據系統時間戳初始化Random
	for len(vals) > 0 {                              //根據牌面陣列長度遍歷
		n := len(vals)                                          //陣列長度
		randIndex := r.Intn(n)                                  //得到隨機index
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1] //最後一張牌和第randIndex張牌互換
		vals = vals[:n-1]
	}
}

func main() {
	ball := Balls{}
	seed := ball.CreateNew(100000)
	fmt.Println("洗牌前: ", seed)
	ball.Shuffle(seed)
	fmt.Println("洗牌後", seed)
}
