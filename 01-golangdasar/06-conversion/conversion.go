package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)
	var nilai8 int8 = int8(nilai16)
	fmt.Println(nilai32) //32768
	fmt.Println(nilai64) //32768
	fmt.Println(nilai16) //-32768 nilai berbalik karena int64 lebih besar dari int16
	fmt.Println(nilai8)  //0 nilai berbalik karena int16 lebih besar dari int8

	var name string = "Daniel Sipangkar"
	var D int8 = int8(name[0])
	var dString = string(D)

	fmt.Println(D, dString)
}
