/**
 * @author Saulo Costa
 * @Date:   2025-06
 * @Last Modified time: 2025-06
 */

/*
* Esse é o nome do pacote. Ele é o main, ou seja, o pacote principal.
* tudo que estiver dentro da função main será executado.
 */
package main

/**
* Importando o pacote fmt. Podemos importar quantos pacotes quisermos.
* https://pkg.go.dev/fmt
* para importar varios pacotes podemos usar o seguinte formato:
* import (
*  "fmt"
*  "math"
*  "time"
* )
)
*/

// import "fmt"
// ou
import (
	// "fmt"
	// renomeando o pacote
	// "fmt" -> "f"
	f "fmt"
)

func main() {
	/* Imprimindo Hello World. O println imprime uma linha. O fmt tem muitas funções:
	* https://golang.org/pkg/fmt */
	// fmt.Println("Hello World")
	f.Println("Hello World")
}

// Para executar o código basta rodar o comando: go run 1-hello-word.go
// Para compilar o código basta rodar o comando: go build 1-hello-word.go
