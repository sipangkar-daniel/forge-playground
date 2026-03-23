package main

import "fmt"

func main() {
	/*
		boolean operation adalah operasi yang menghasilkan nilai boolean

		ada 3 operator &&, || dan !

		&& akan bernilai true jika kedua operasi di sebelah kiri dan kanan adalah true
		|| akan bernilai true jika salah satu operasi di sebelah kiri dan kanan adalah true
		! akan membalikkan nilai boolean

		true && true -> true
		true && false -> false
		false && true -> false
		false && false -> false

		true || true -> true
		true || false -> true
		false || true -> true
		false || false -> false

		!true -> false
		!false -> true
	*/

	nilaiAkhir := 80
	absensi := 3

	result := nilaiAkhir >= 80 && absensi < 3

	if result {
		fmt.Println("You passed")
	} else {
		fmt.Println("You failed")
	}

}
