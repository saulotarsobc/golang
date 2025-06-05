package main

import "fmt"

/* Funcões
- Você pode criar funções para reutilizar o código.
- Isso facilita a manutenção do código.
*/

func main() {
	fmt.Println(soma(1, 2))
}

/* Soma
- Recebe como parâmetro dois números de tipo int.
- Retorna a soma de dois números de tipo int.
- Você pode 'tipar' o retorno da função.
*/
func soma(a int, b int) int {
	return a + b
}

// https://youtu.be/eTHGCR2CUeA?list=PLIIX-IKjIiwOpAr_kyvpxTVyvUoxXqGEQ&t=261
