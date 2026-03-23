package main

import "fmt"

func main() {
	/**
	operator matematika
	+ - / * %

	augmented assigment
	+= -= /= *= %=
	*/

	var a int = 10
	var b int = 20

	var c int = a + b
	d := c % b

	fmt.Println(a, "+", b, "=", c, "%", b, "=", d)
}
