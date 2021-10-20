package main

import "fmt"

func main() {
	var n int
	fmt.Println("Jawaban No 8 : Membuat tabel perkalian")
	fmt.Println("=")
	fmt.Print("Masukkan bilangan : ")
	fmt.Scanf("%d", &n)
	fmt.Println("=")
	multiply_table(n)
}

func multiply_table(n int) {
	if n > 0 && n <= 30 {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				fmt.Print(i*j, " ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Bilangan yang diperbolehkan hanya 1 - 30")
	}
}
