package main

import (
	"fmt"
)

var base64Alphabet []rune

func main() {
	// create array that maps 6 bit binary chunks to characters
	for l := 'A'; l <= 'Z'; l++ {
		base64Alphabet = append(base64Alphabet, l)
	}
	for l := 'a'; l <= 'z'; l++ {
		base64Alphabet = append(base64Alphabet, l)
	}
	for l := '0'; l <= '9'; l++ {
		base64Alphabet = append(base64Alphabet, l)
	}
	base64Alphabet = append(base64Alphabet, '+', '/')

	fmt.Println(encode([]byte("hel")))
}

func encode(bytes []byte) string {
	sixBitChunks := []byte{}
	sixBitChunk := byte(0)
	bitsAddedToChunk := 0
	for _, b := range bytes {
		for i := 7; i >= 0; i-- {
			bit := getBit(b, i)
			sixBitChunk += bit
			sixBitChunk <<= 1

			bitsAddedToChunk += 1
			if bitsAddedToChunk == 6 {
				sixBitChunk >>= 1
				sixBitChunks = append(sixBitChunks, sixBitChunk)
				sixBitChunk = 0
				bitsAddedToChunk = 0
			}
		}
	}

	result := []rune{}
	for _, chunk := range sixBitChunks {
		result = append(result, base64Alphabet[chunk])
	}
	return string(result)
}

func getBit(b byte, pos int) byte {
	if (b & (1 << pos)) != 0 {
		return 1
	} else {
		return 0
	}
}
