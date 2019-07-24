package codingame

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

var glyphs5x4 = Dimensions{Height: 5, Width: 4}

var glyphs11x20 = Dimensions{Height: 11, Width: 20}

func TestAsciiArt(t *testing.T) {
	tests := []struct {
		name       string
		dimensions Dimensions
		reqText    string
		expected   string
	}{
		{
			name:       "one letter - E",
			dimensions: glyphs5x4,
			reqText:    "E",
			expected:   "asciiE.txt",
		},
		{
			name:       "MANHATTAN",
			dimensions: glyphs5x4,
			reqText:    "MANHATTAN",
			expected:   "asciiManhattan.txt",
		},
		{
			name:       "Manhattan mixed case",
			dimensions: glyphs5x4,
			reqText:    "ManhAtTan",
			expected:   "asciiManhattan.txt",
		},
		{
			name:       "Manhattan non-alpha",
			dimensions: glyphs5x4,
			reqText:    "M@NH@TT@N",
			expected:   "asciiManhattanWithAt.txt",
		},
		{
			name:       "Manhattan with another ASCII representation",
			dimensions: glyphs11x20,
			reqText:    "MANHATTAN",
			expected:   "asciiManhattan11x20.txt",
		},
		{
			name:       "go-l1ng - non-alpha 11x20",
			dimensions: glyphs11x20,
			reqText:    "go-l1ng",
			expected:   "asciiGolangNonAlpha11x20.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out strings.Builder
			glyphs := readGlyphs(tt.dimensions, t)

			// fmt.Fprintf(
			// 	os.Stderr,
			// 	"%s: dimensions=%#v, reqText=%q, glyphs=\n%%q\n",
			// 	tt.name, tt.dimensions, tt.reqText, //glyphs,
			// )

			AsciiArt(tt.dimensions, glyphs, tt.reqText, &out)

			expectedStr := readTestData(tt.expected, t)
			if out.String() != expectedStr {
				t.Errorf(
					"AsciiArt(..., %#v, ...)\n!!GOT:\n%s\n!!EXPECTED:\n%s\n",
					tt.reqText, out.String(), expectedStr,
				)
			}
		})
	}
}

// this doesn't work because of https://github.com/golang/go/issues/26460
// func ExampleAsciiArt_e_capcase() {
// 	AsciiArt(glyphs4x5.height, glyphs4x5.length, glyphs4x5.glyphs, "E", os.Stdout)

// 	/*Output:*/
// 	/*### */
// 	/*#   */
// 	/*##  */
// 	/*#   */
// 	/*###*/
// }

func TestIsAsciiLetter(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"alpha-lower", args{'a'}, true},
		{"alpha-upper", args{'D'}, true},
		{"symbol-hyphen", args{'-'}, false},
		{"symbol-dot", args{'.'}, false},
		{"symbol-question-mark", args{'?'}, false},
		{"numeric", args{'2'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAsciiLetter(tt.args.char); got != tt.want {
				t.Errorf("IsAsciiLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*  Read in test data file in `testdata/` directory, return contents as string
 */
func readTestData(filename string, t testing.TB) string {
	datafile := filepath.Join("testdata", filename)
	bytes, err := ioutil.ReadFile(datafile)
	if err != nil {
		t.Fatal(err)
	}
	return string(bytes)
}

/* Load glyphs from test data file, return list of lines
 */
func readGlyphs(glyphDim Dimensions, t testing.TB) []string {
	contents := readTestData(
		fmt.Sprintf("glyphs%dx%d.txt", glyphDim.Height, glyphDim.Width), t,
	)
	return strings.Split(contents, "\n")
}
