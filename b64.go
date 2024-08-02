package b64

import (
	"slices"
)

var base64Alphabet []byte

func init() {
	// create array that maps 6 bit binary chunks to characters
	for l := 'A'; l <= 'Z'; l++ {
		base64Alphabet = append(base64Alphabet, byte(l))
	}
	for l := 'a'; l <= 'z'; l++ {
		base64Alphabet = append(base64Alphabet, byte(l))
	}
	for l := '0'; l <= '9'; l++ {
		base64Alphabet = append(base64Alphabet, byte(l))
	}
	base64Alphabet = append(base64Alphabet, byte('+'), byte('/'))
}

func Encode(bytes []byte) string {
	// number of zeroed bytes to add to ensure bytes is a multiple of 3
	bytesToAdd := (3 - len(bytes)%3) % 3
	// add those zeroed bytes
	for i := 0; i < bytesToAdd; i++ {
		bytes = append(bytes, byte(0))
	}

	// iterate over the bytes => iterate over each bit in byte => build up six bit chunks
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

	// convert six bit chunks to chars
	result := []byte{}
	for _, chunk := range sixBitChunks {
		result = append(result, base64Alphabet[chunk])
	}

	// convert any zeroed bytes that were added to end to '='
	for i := len(result) - bytesToAdd; i < len(result); i++ {
		result[i] = '='
	}

	return string(result)
}

func Decode(str string) []byte {
	bytes := []byte{}
	byteChunk := byte(0)
	bitsAddedToChunk := 0
	// iterate over the characters => get six bits that each character represents => iterate over bits => build up byte chunks
	for i, _ := range str {
		c := str[i]
		if c == '=' {
			continue
		}
		sixBitChunk := byte(slices.Index(base64Alphabet, c))
		for i := 5; i >= 0; i-- {
			bit := getBit(sixBitChunk, i)
			byteChunk += bit
			byteChunk <<= 1

			bitsAddedToChunk += 1
			if bitsAddedToChunk == 8 {
				byteChunk >>= 1
				bytes = append(bytes, byteChunk)
				byteChunk = 0
				bitsAddedToChunk = 0
			}
		}
	}

	return bytes
}

func getBit(b byte, pos int) byte {
	if (b & (1 << pos)) != 0 {
		return 1
	} else {
		return 0
	}
}
