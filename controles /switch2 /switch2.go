package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // Isso seria um Switch true. Ele vai procurar o primeiro case que for true
	case t.Hour() < 12:
		fmt.Println("Bom dia !")
	case t.Hour() < 18:
		fmt.Println("Boa tarde !")
	default:
		fmt.Println("Boa noite !")

	}
}
