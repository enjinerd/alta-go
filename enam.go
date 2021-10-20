package main

import "fmt"

func main() {
	fmt.Println("Jawaban No 6 : Hasil perpangkatan")
	fmt.Println("=")
	var x, n int
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scanf("%d", &x)
	fmt.Print("Masukkan Pangkat : ")
	fmt.Scanf("%d", &n)
	fmt.Println("=")
	fmt.Println("Hasilnya : ", exponentiation(x, n))
}

func exponentiation(x, n int) int {
	res := 1

	if n == 0 {
		return 1
	}

	for i := 0; i < n; i++ {
		res = res * x
	}

	return res
}
