package main

import (
	"fmt"
	"time"
)

func main() {

	var count int = 50

	start := time.Now()
	fmt.Println(fibRecursive(count))
	elapsed := time.Now().Sub(start)
	fmt.Println("took " + elapsed.String())

	start = time.Now()
	fmt.Println(fibFast(count))
	elapsed = time.Now().Sub(start)
	fmt.Println("took " + elapsed.String())

}

func fibRecursive(seq int) int {
	if seq <= 1 {
		return seq
	}
	return fibRecursive(seq-1) + fibRecursive(seq-2)
}

func fibFast(seq int) int {
	if seq <= 1 {
		return seq
	}

	s := make([]int, seq)
	s[0] = 0
	s[1] = 1
	for i := 2; i < seq; i++ {
		s[i] = s[i-1] + s[i-2]
	}

	return s[seq-1] + s[seq-2]
}
