package main

import "fmt"

/*
array
-> kumpulan data dengan tipe data yang sama
-> ukuran/size array ditentukan sebelum array dibuat, ukuran tidak dapat ditambah/dikurang setelah dibuat
-> array memiliki index untuk mengakses data dalam array
*/

func main() {

	var names [5]string
	names[0] = "Daniel"
	names[1] = "Sipangkar"
	names[2] = "Rahmat"
	names[3] = "Saputra"

	fmt.Println(names)

	var age [5]int
	age[0] = 20
	age[1] = 25
	age[2] = 30
	age[3] = 35
	age[4] = 40

	var addresses = [5]string{"Jl", "Jl", "Jl", "Jl", "Jl"}
	fmt.Println(addresses)

	//akses array dengan index, index array dimulai dari 0

	fmt.Println("My name is", names[0])

	/*
		function dalam array
		len(array) -> menghitung ukuran array
		cap(array) -> menghitung kapasitas array
		array[index] -> mengakses data array
		array[index] = value -> mengubah data array
	*/

	panjangArray := len(names)
	kapasitasArray := cap(names)
	dataPertama := names[0]
	fmt.Printf("Panjang array = %d, Kapasitas array = %d, Data pertama = %s \n", panjangArray, kapasitasArray, dataPertama)

	names[0] = "Joko"
	fmt.Println("nama index pertama diganti menjadi", names[0])

}
