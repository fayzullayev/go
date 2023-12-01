package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (s MyString) FindVowels() []rune {
	var vowels []rune

	for _, r := range s {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowels = append(vowels, r)
		}
	}

	return vowels
}

func main() {

	//name := MyString("Mirodil Fayzullayev")
	//
	//var v VowelsFinder
	//
	//v = name
	//
	//fmt.Printf("Vowels are %c\n", v.FindVowels())

	//fmt.Printf("%v, %x, %08b, %c --- %08b\n", '💕', '💕', '💕', '💕', "💕")

	str := "💕"
	StringToBinary(str)
	//fmt.Println("Binary representation:", binary)

}

func StringToBinary(s string) {

	for _, c := range s {
		fmt.Println(1)
		fmt.Printf("%08b\n", c)
	}

}
