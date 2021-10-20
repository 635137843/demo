package main

import "fmt"

func main() {
	list := search([]int{3, 4, 8, 1, 2, 7, 9, 5, 6})
	fmt.Println("切片数据：", list)
}
func findSmall(list []int) int {
	small := list[0]
	smallIndex := 0
	for k, v := range list {
		if v < small {
			small = v
			smallIndex = k
		}
	}
	return smallIndex
}
func search(list []int) []int {
	var result []int
	maxIndex := len(list)
	for {
		smallIndex := findSmall(list)
		result = append(result, list[smallIndex])
		list = append(list[:smallIndex], list[smallIndex+1:]...)
		if len(result) == maxIndex {
			fmt.Println("排序完毕！")
			return result
		}
	}
}
