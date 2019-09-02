package main

import "fmt"

// quais são os valores que voltam como zero das variaveis basicas.
func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int // isso é um ponteiro 
	fmt.Println("%v %v %v %s %v", a, b, c, d, e)
}
