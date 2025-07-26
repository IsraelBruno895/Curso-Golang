package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) // forma mais verbosa

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste" + string(97)) // será inpresso na concatenação o equivalente de 97 na tabela ASCII

	// int para string
	fmt.Println("Teste" + strconv.Itoa(123))

	num, _ := strconv.Atoi("123") // a função retorna um número e um erro com o underline estou ignorando o erro
	fmt.Println(num - 122)

	// string para bool

	b, _ := strconv.ParseBool("1") // converte o valor para booleano
	if b {
		fmt.Println("verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}
