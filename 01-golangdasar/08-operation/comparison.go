package main

import "fmt"

func main() {
	/**
	comparison operator menghasilkan nilai boolean untuk membandingkan data

	> lebih besar
	< lebih kecil
	>= lebih besar sama dengan
	<= lebih kecil sama dengan
	== sama dengan
	!= tidak sama dengan

	*/

	value := 78

	result := value >= 80

	if result {
		fmt.Println("You passed")
	} else {
		fmt.Println("You failed")
	}

}
