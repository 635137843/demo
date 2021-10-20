package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		list = append(list, i+1)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		v := rand.Intn(1_000_000-1) + 1
		fmt.Printf("针对%d进行二分查找：\n", v)
		value := search(list, v)
		fmt.Println("猜的值为：", value)
	}
}
func search(list []int, guess int) int {
	low := 0
	high := len(list) - 1

	step := 0
	for {
		step++
		if low <= high {
			mid := (low + high) / 2
			if mid == guess {
				fmt.Printf("总共查询了%d次\n", step)
				return mid
			}
			if mid < guess {
				low = mid + 1
			}
			if mid > guess {
				high = mid - 1
			}
		}
	}
}
