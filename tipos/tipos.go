package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 100)
	fmt.Println("Literal intrito é", reflect.TypeOf(32000)) // função reflect.TypeOf() retorna o tipo do Literal

	// sem sinal (só possitivos)... uint8 uint16 uint32 uint64

	var b byte = 255 // byte é um apelido para o uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64

	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é ", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do loteral 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "ola meu nome é Israel"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1)) // len() devolve o tamanho da string

	// string com multiplas linhas
	// em Go é possivel criar strings com o abrir e fechar crase
	s2 := `Ola  
	meu
	nome
	é
	Israel`
	fmt.Println("O tamanho da string é", len(s2))
	fmt.Println(s2)

	// char???
	// var x char = 'b' não  funciona é um int32

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
