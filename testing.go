package main

import (
	"fmt"
)

type Cipher interface {
	Encrypt() interface{}
}

type myString string

type myArray []int

func (s myString) Encrypt() interface{} {
	letterMapping := "UMRSQPBOLEXTZYAKJVCNHDWGIF"
	result := ""
	for _, char := range s {
		if 'A' <= char && char <= 'Z' {
			index := char - 'A'
			result += string(letterMapping[index])
		} else {
			result += string(char)
		}
	}
	return result
}

func (ia myArray) Encrypt() interface{} {
	result := make(myArray, len(ia))
	for i, num := range ia {
		if num%2 == 0 {
			result[i] = num / 2
		} else {
			result[i] = 3*num + 1
		}
	}
	return result
}

func printCipher(c Cipher) {
	fmt.Println(c.Encrypt())
}

func main() {

	var input myString
	fmt.Println("Enter String:")
	fmt.Scanln(&input)
	printCipher(input)

	fmt.Println("Enter size of array:")
	var size int
	fmt.Scanln(&size)

	var intArray myArray
	fmt.Println("Enter array elements:")
	if size > 0 {
		for i := 0; i < size; i++ {
			var element int
			fmt.Scan(&element)
			intArray = append(intArray, element)
		}
		printCipher(intArray)
	} else {
		fmt.Println("Enter size>0")
	}

}
