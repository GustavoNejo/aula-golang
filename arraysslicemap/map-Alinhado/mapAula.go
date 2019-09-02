package main

import "fmt"

func amain() {
	n := make(map[int]string)

	n[10] = "A"
	n[9] = "a\n"
	fmt.Println(n)
}
