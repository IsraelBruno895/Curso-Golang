package main

import "fmt"

func main() {
	// Go não tem aritmética de ponteiro

	i := 1

	var p *int = nil

	p = &i // pega o endereço da variável
	*p++   // desreferenciabdo (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)

}
