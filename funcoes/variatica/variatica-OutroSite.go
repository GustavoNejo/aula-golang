package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("nums", nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
}
func main() {

}
