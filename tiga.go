package main

import "fmt"

func main() {
	fmt.Println("Jawaban No 3 : Faktor Bilangan")
	fmt.Println("=")
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Faktor bilangannya adalah: ")
	factor(n)
}

func factor(n int) {
	if n <= 1 {
		fmt.Print("Inputan harus bilangan bulat")
	}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}
