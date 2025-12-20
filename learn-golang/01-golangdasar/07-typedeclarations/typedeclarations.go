package main

import "fmt"

func main() {
	//type or alias

	type Name string

	var name Name = "Daniel Sipangkar"

	fmt.Println(name)
	fmt.Println(Name("Sipangkar"))

}
