package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Println("Linha")

	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 3.141516

	//fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x) // Sprint retorna uma string do valor passado

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)
	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d) // %v  é um marcador de formato generico serve para todos os tipos

}
