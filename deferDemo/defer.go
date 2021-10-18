package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tryDefer()
}

func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
	writeFile("file.txt")
}

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	write := bufio.NewWriter(file)
	defer func(write *bufio.Writer) {
		err := write.Flush()
		if err != nil {
			fmt.Println("写入失败！")
		}
	}(write)
	for i := 0; i < 20; i++ {
		_, err := fmt.Fprintln(write, strconv.Itoa(i))
		if err != nil {
			return
		}
	}
}
