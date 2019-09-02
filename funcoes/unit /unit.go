package main

import "fmt"

func init() { //A func init pode ser usada em qualquer arquivo para inicilizar alguma func
	// a func init ela é exibida em qualquer lugar quando solicitada pelo terminal 
	fmt.Println("Inicializando...")
}
func main() {
	fmt.Println("Main")
}
