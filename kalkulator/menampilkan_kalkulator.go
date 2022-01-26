package main

import "fmt"

func main() {
	var operan1 int
	var operan2 int
	var operator string

	fmt.Println("masukan operan1")
	fmt.Scanln(&operan1)
	fmt.Println("masukan operator (+,-,*,/)")
	fmt.Scanln(&operator)
	fmt.Println("masukan operan2")
	fmt.Scanln(&operan2)

	result := 0
	//operator penjumlahan
	if operator == "+" {
		result = operan1 + operan2

		//cetak hasil
		fmt.Println("result", operan1, operator, operan2, "=", result)
	}
}
