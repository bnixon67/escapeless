package main

import (
	"crypto/rand"
	"fmt"
)

func Encode(src []byte) {
	chunkSize := 255

	used := make([]int8, 256)

	// mark all byte values used in chunk
	for n := 0; n < chunkSize; n++ {
		used[src[n]] = 1
	}

	// print used
	for n := 0; n < len(used); n = n + 32 {
		fmt.Println(used[n : n+32])
	}

	// find greatest byte value not used within chunk
	substitute := 0
	for n := 255; n >= 0; n-- {
		if used[n] == 0 {
			substitute = n
			break
		}
	}

	fmt.Println("substitute value is", substitute)
}

func main() {
	src := make([]byte, 1024)

	//rand.Read(src)

	for n := 0; n < len(src); n++ {
		src[n] = byte(n)
	}

	rand.Read(src)

	Encode(src)

	//fmt.Println(src)
}
