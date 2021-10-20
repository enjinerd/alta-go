package main

import "fmt"

func main() {
	var rows int
	fmt.Println("Jawaban No 7 : Membuat segitiga asterik")
	fmt.Println("=")
	fmt.Print("Masukkan jumlah tinggi segitiga (kolom): ")
	fmt.Scanf("%d", &rows)
	asterik(rows)
}

func asterik(rows int) {
	j := 0
	for i := 1; i <= rows; i++ {
		j = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print(" ")
		}
		for {
			fmt.Print("* ")
			j++
			if j == i {
				break
			}
		}
		fmt.Println("")
	}
}
