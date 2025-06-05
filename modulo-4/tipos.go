package main

import "fmt"

/* Tipos
bool - true/false
string - sequence of bytes
int
float
*/

func main() {
	// Usando o Printf do package fmt
	// %T - type - Server para mostrar o tipo de uma variável
	// %v - value - Server para mostrar o valor de uma variável

	// bool é um valor booleano. Ou seja, true ou false.
	fmt.Printf("Type: %T Value: %v\n", true, true)

	// string é uma sequência de bytes
	fmt.Printf("Type: %T Value: %v\n", "saulo", "saulo")

	// int é um inteiro. Ou seja, um número inteiro. Ex: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	// Se passarmos o número com aspas, ele será considerado uma string. Ex: "1"
	fmt.Printf("Type: %T Value: %v\n", 1, 1)
	fmt.Printf("Type: %T Value: %v\n", "1", "1")

	// float e um número real. Ou seja, um número com ponto flutuante. Ex: 1.1
	// Existem do tipo float32 e float64.
	// float32 usa 32 bits e float64 usa 64 bits.
	fmt.Printf("Type: %T Value: %v\n", 1.1, 1.1)
}
