package main

import "fmt"

func main() {
	var r, T float64
	fmt.Println("Jawaban No 1 : Menghitung Luas Permukaan Tabung")
	fmt.Println("=")
	fmt.Print("Masukkan luas alas (cm): ")
	fmt.Scanf("%f", &r)
	fmt.Print("Masukkan luas selimut (cm): ")
	fmt.Scanf("%f", &T)
	fmt.Println("Maka luas permukaannya : ", tube(r, T), " cm²")
}

func tube(r, T float64) float64 {
	const pi float64 = 3.14
	return 2.0 * pi * r * (r + T)
}
