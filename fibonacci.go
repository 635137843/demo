package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() fiGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type fiGen func() int

func (f fiGen) Read(p []byte) (n int, err error) {
	s := f()
	if s > 20000 {
		return 0, io.EOF
	}
	v := fmt.Sprintf("%d\n", s)
	return strings.NewReader(v).Read(p)
}

func readFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	readFileContent(f)
}
