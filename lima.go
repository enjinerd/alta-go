package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Jawaban No 5 : Mengecek jenis kata apakah palindrome")
	fmt.Println("=")
	fmt.Print("Masukkan kata : ")
	fmt.Scanf("%s", &s)
	fmt.Println("=")
	fmt.Println("Kata : ", s, ", Merupakan : ", palindrome(s))
}

func palindrome(s string) string {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return "Bukan palindrome"
		}
	}
	return "Palindrome"
}
