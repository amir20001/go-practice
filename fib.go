package main

import (
	"fmt"
	"time"
)

var cache [100]int

func main() {

	var count int = 46

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
		cache[seq] = seq
		return seq
	}

	if cache[seq] != 0 {
		return cache[seq]
	}

	value := fibFast(seq-1) + fibFast(seq-2)
	cache[seq] = value
	return value
}
