package main

import "fmt"

func f1() { // func 1 nao recebe nd de dado e retorna nd de dado
	fmt.Println("F1")
}

func f2(p1 string, p2 string) { // func 2 recebe dois parametros, mas ela nao retorna nd
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string { // func 3 nao recebe nd, mas ela retorna um parametro (Uma string no caso)
	return "F3" // toda vez que eu pedir para retornar um string eu tenho que fazer isso em todo o bloco, se nao pode gerar um erro
}

func f4(p1, p2 string) string { // func 4 ele recebe dois parametro e retorna um inico parametro
	// eu tambem posso escrever primeiro os parametros e dps mostrar qual o tipo do parametro
	return fmt.Sprintf("F4: %s %s", p1, p2) // Obrigatoriamente por retornar uma string eu tenho que fazer um return
	// a func Sprintf nao imprime no console, mas sim retorna uma string formatada
}
func f5() (string, string) { // essa func ela nao recebe nenhum parametro, mas ela retorna dois parametros
	return "Retorno 1", "Retorno 2"
}

func main() {
	// abaixo é somente chamadas de cada uma das funcões
	f1()
	f2("param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5", r51, r52)
}
