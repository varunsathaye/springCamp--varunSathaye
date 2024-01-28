package main

import (
	"fmt"
)

type Cipher interface {
	Encrypt() interface{}
}

type myString string

type myArray []int

type myMap map[rune]int

func (s myString) Encrypt() interface{} {
	var result []int
	for _, char := range s {

		if 'A' <= char && char <= 'Z' {

			result = append(result, int(char-64))
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
func (m myMap) Encrypt() interface{} {
	var result []int
	for key, value := range m {
		result = append(result, int(key)+value)
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
	var m myMap
	var size2 int
	fmt.Println("Enter number of entries in map")
	fmt.Scanln(size2)
	fmt.Println("Enter key-value pairs for the map (e.g., 'a 10 b 20'):")
	m = make(myMap)

	for i := 0; i < size2; i++ {
		var key rune
		var value int
		fmt.Scan(&key, &value)
		m[key] = value
	}
	printCipher(m)

}
