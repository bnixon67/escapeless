/*
Copyright 2021 Bill Nixon

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
