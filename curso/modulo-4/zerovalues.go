package main

import "fmt"

/* zerovalues
Uma variável declarada sem um valor inicial recebe um zero value.
Ou seja, uma variável declarada sem um valor inicial recebe um zero value.
Cada tipo de variável tem um zero value diferente. Ex:

int -> 0
string -> ""
bool -> false
*/
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Inteiro:%v\n", i)
	fmt.Printf("Float:%v\n", f)
	fmt.Printf("Bool:%v\n", b)
	fmt.Printf("String:%v\n", s)
}
