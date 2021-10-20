package main

import "fmt"

func main() {
	var score int
	var name string
	fmt.Println("Jawaban No 2 : Mefilter Grade Nilai")
	fmt.Println("=")
	fmt.Print("Masukkan Nama Siswa : ")
	fmt.Scanf("%s", &name)
	fmt.Print("Masukkan Nilai Siswa : ")
	fmt.Scanf("%d", &score)
	fmt.Println("=")
	fmt.Println("Nilai : ", grade(score))
}

func grade(score int) string {
	if score < 0 || score > 100 {
		return "Nilai invalid, nilai yang diterima : 0 - 100"
	} else {
		switch {
		case score < 35:
			return "E"
		case score < 50:
			return "D"
		case score < 65:
			return "C"
		case score < 80:
			return "B"
		default:
			return "A"
		}
	}
}
