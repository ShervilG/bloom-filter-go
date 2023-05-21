package main

import (
	"fmt"
	"os"
	"strings"
)

var keySetSize int = 31
var bloomFilter int = 0

func saveToBloomFilter(key string) {
	charSum := 0
	for _, char := range key {
		charSum += int(char)
	}

	filterIndex := charSum % keySetSize
	bloomFilter = bloomFilter | 1<<(filterIndex+1)
}

func keyExists(key string) bool {
	charSum := 0
	for _, char := range key {
		charSum += int(char)
	}

	filterIndex := charSum % keySetSize
	tempFilter := bloomFilter
	tempFilter = tempFilter >> (filterIndex + 1)

	return (tempFilter)&1 == 1
}

func main() {
	println("Welcome to Bloom Filter example")

	for true {
		println("1. Press 'Q' to quit")
		println("2. Press 'E' to enter element")
		println("2. Press 'C' to check element")

		var firstInput string
		var secondInput string
		fmt.Scan(&firstInput)

		switch strings.ToUpper(firstInput) {
		case "Q":
			os.Exit(500)

			break
		case "E":
			print("Enter key: ")
			fmt.Scan(&secondInput)

			saveToBloomFilter(secondInput)

			break
		case "R":
			print("Key exists ? : ")
			fmt.Scan(&secondInput)

			if keyExists(secondInput) {
				println("Maybe..")
			} else {
				println("Not at all !")
			}

			break
		}
	}
}
