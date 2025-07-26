package main

import (
	"fmt"
	"time"
)

func main() {
	d1 := time.Unix(0, 0) // zero horas e zero segunds da época unix de 1 de janeiro de 1970
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2)) // usado paracomparar dois objetos no tempo de forma confiáel o metodo Equal() torna a comparação de pontos no tempo mais robusta e precisa

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas:", p1 == p2)

}
