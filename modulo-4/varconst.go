package main

import "fmt"

func main() {
	// Variaveis
	var nome string
	var altura float64
	// nome = 1 // Vai dar erro, pois a variável nome foi declarada como string
	nome = "Saulo"
	altura = 1.80
	fmt.Println(nome, altura)

	// Inferencia de tipos automática
	var sobrenome = "Costa" // para o go essa variável tem o tipo string
	var idade = 20          // para o go essa variável tem o tipo int
	fmt.Println(sobrenome, idade)

	// Declarar mais de uma variável com o mesmo tipo
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// Declarar mais de uma variável com o mesmo tipo
	var (
		d = 1
		e = 2
		f = 3
	)
	fmt.Println(d, e, f)

	// Declatar variável com ':='. Significa que ela foi declarada e iniciada na mesma linha
	numero := 1
	numeros := []int{1, 2, 3}
	fmt.Println(numero, numeros)

	// Constantes. Ou seja uma variável que nao pode ter seu valor alterado
	const pi = 3.14
	const pi2 float64 = 3.14
	fmt.Println(pi, pi2)
}
