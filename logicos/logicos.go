package main

import "fmt"

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2 // ou exclusivo
	comprarSorvete := trabalho1 || trabalho2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, Sorvete := compras(true, true)
	fmt.Printf("TV50: %t, TV32: %t, Sorvete: %t", tv50, tv32, !Sorvete)
}
