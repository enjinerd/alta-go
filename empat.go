package main

import "fmt"

func main() {
	fmt.Println("Jawaban No 4 : Mengecek bilangan prima")
	fmt.Println("=")
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanf("%d", &n)
	fmt.Println("=")
	fmt.Println("Bilangan : ", n, ", Merupakan : ", primeNumber(n))
}

func primeNumber(n int) string {
	if n <= 1 {
		return "Bukan bilangan prima"
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return "Bukan bilangan prima"
		}
	}
	return "Bilangan Prima"
}
