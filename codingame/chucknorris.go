package codingame

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type ChuckEncoder []string

// Write a type block for `bit` followed by `count * "0"` to the encoder
func (this *ChuckEncoder) Encode(bit rune, count int) {
	typeBlock := typeBlockFor(bit)
	encBlock := strings.Repeat("0", count)
	// fmt.Fprintf(
	// 	os.Stderr,
	// 	"Encode(%q, %d): typeBlock=%q, encBlock=%q\n",
	// 	bit, count, typeBlock, encBlock,
	// )
	*this = append(*this, typeBlock, encBlock)
}

// Return all encoded blocks as space-separated string
func (this *ChuckEncoder) String() string {
	return strings.Join([]string(*this), " ")
}

/* Encode `message` using Chuck Norris’ technique
 */
func ChuckNorris(message string) string {
	var sb strings.Builder
	for _, char := range message {
		fmt.Fprintf(&sb, "%07b", char)
	}

	bits := sb.String()
	sb.Reset()
	// fmt.Fprintf(os.Stderr, "message=%q, bits=%q\n", message, bits)
	var encodedBlocks ChuckEncoder
	lastBit := utf8.RuneError
	seriesCounter := 1
	finalIx := utf8.RuneCountInString(bits) - 1
	for ix, bit := range bits {
		if bit == lastBit {
			seriesCounter++
		}
		// fmt.Fprintf(os.Stderr, "bit=%q, lastBit=%q, counter=%d, ix=%d/%d\n", bit, lastBit, seriesCounter, ix, finalIx)
		if lastBit != utf8.RuneError && (bit != lastBit || ix == finalIx) {
			encodedBlocks.Encode(lastBit, seriesCounter)
		}
		if bit != lastBit {
			// reset counter
			seriesCounter = 1
			if ix == finalIx {
				// also write this bit as it's the last iteration
				encodedBlocks.Encode(bit, seriesCounter)
			}
		}
		lastBit = bit
	}
	return encodedBlocks.String()
}

func typeBlockFor(bit rune) string {
	if bit == '0' {
		return "00"
	}
	return "0"
}
