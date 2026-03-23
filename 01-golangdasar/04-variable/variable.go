package main

import "fmt"

func main() {
	var name string

	fmt.Println(name)

	name = "Daniel Sipangkar"
	fmt.Println("My name is", name)

	age := 20 //simplify var at first declaration
	age = 25
	fmt.Println("age", age, "years old")

	institute := "Institut Teknologi Sumatera"
	fmt.Println("I study at", institute)

	// multiple var declaration
	var (
		office         = "PT. Bank Mandiri"
		position       = "Software Engineer"
		yoe            = 2021
		salary         = 99999999
		salaryCurrency = "Rp"
	)

	fmt.Printf("Saya bekerja di %s sebagai %s sejak %d dengan salary %s%d \n", office, position, yoe, salaryCurrency, salary)
}
