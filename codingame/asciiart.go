package codingame

import "fmt"
import "io"
import "os"
import "bufio"
import "unicode"

type Dimensions struct {
	Height int
	Width  int
}

func CgRunAsciiArt() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var glyphWidth int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &glyphWidth)

	var glyphHeight int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &glyphHeight)

	glyphDimensions := Dimensions{glyphHeight, glyphWidth}

	scanner.Scan()
	reqText := scanner.Text()

	// fmt.Fprintf(
	// 	os.Stderr,
	// 	"glyphDimensions: %#v, reqText: %s\n",
	// 	glyphDimensions,
	// 	reqText,
	// )

	var glyphs = make([]string, glyphHeight)
	for i := 0; i < glyphHeight; i++ {
		scanner.Scan()
		row := scanner.Text()
		// fmt.Fprintf(os.Stderr, "row: %q\n", row)
		glyphs[i] = row
	}

	AsciiArt(glyphDimensions, glyphs, reqText, os.Stdout)
}

func AsciiArt(glyphDimensions Dimensions, glyphs []string, reqText string, out io.Writer) {
	// fmt.Fprintln(os.Stderr, "text:")
	var reqGlyphIndices = make([]int, len(reqText))
	for i, char := range reqText {
		var glyphOffset int
		if IsAsciiLetter(char) {
			char = unicode.ToUpper(char)
			glyphOffset = int(char - 'A')
		} else {
			char = '?'
			glyphOffset = 26 // ? is the last glyph, after all letters
		}
		glyphOffsetIx := glyphOffset * glyphDimensions.Width
		fmt.Fprintf(
			os.Stderr,
			"char=%q, offset=%d, index=%d\n",
			char, glyphOffset, glyphOffsetIx,
		)
		reqGlyphIndices[i] = glyphOffsetIx
	}

	for _, row := range glyphs {
		if len(row) == 0 {
			continue
		}
		for _, glyphIx := range reqGlyphIndices {
			// fmt.Fprintf(
			// 	os.Stderr,
			// 	"glyphIx=%d, glyphWidth=%d, row=%q\n",
			// 	glyphIx, glyphDimensions.Width, row,
			// )
			glyphRow := row[glyphIx : glyphIx+glyphDimensions.Width]
			// fmt.Fprintf(os.Stderr, "%q", glyphRow)
			fmt.Fprint(out, glyphRow)
		}
		fmt.Fprintln(out)
	}
}

func IsAscii(char rune) bool {
	return char <= unicode.MaxASCII
}

func IsAsciiLetter(char rune) bool {
	return IsAscii(char) && unicode.IsLetter(char)
}
