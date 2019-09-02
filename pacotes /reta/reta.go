package main

import "math"

// iniciando um pacote com letra maiúscula ele se tornará PUBLICO (visivel dentro e fora do pacote)
// iniciando um pacote com letra minuscula ele se tornará PRIVADO (visivel apenas dentro do pacote)

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) { // essa func sera visivel apenas dentro da func
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia sera visivel dentro e fora da chave, pq o D esta em maiusculo
//distancia é resposavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
 