package aritmetica

//Pi é uma visível fora do pacote pq começa com letra maiúscula
const Pi = 3.1416 

//Soma é pública
func Soma(v1, v2 int) int {
	return v1 + v2
}

//Privado pq começa com _
func _Subtracao(v1, v2 int) int {
	return v1 - v2
}

//Privado pq começa com letra minúscula
func multiplicacao(v1, v2 int) int {
	return v1 * v2
}