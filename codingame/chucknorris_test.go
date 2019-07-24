package codingame

import (
	"testing"
)

func TestChuckNorris(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
	}{
		{"upper-c", "C", "0 0 00 0000 0 00"},
		{"upper-cc", "CC", "0 0 00 0000 0 000 00 0000 0 00"},
		{"percent", "%", "00 0 0 0 00 00 0 0 00 0 0 0"}, // 0100101
		{
			"message-from-chuck",
			"Chuck Norris' keyboard has 2 keys: 0 and white space.",
			"0 0 00 0000 0 0000 00 0 0 0 00 000 0 000 00 0 0 0 00 0 0 000 00 000 0 0000 00 0 0 0 00 0 0 00 00 0 0 0 00 00000 0 0 00 00 0 000 00 0 0 00 00 0 0 0000000 00 00 0 0 00 0 0 000 00 00 0 0 00 0 0 00 00 0 0 0 00 00 0 0000 00 00 0 00 00 0 0 0 00 00 0 000 00 0 0 0 00 00000 0 00 00 0 0 0 00 0 0 0000 00 00 0 0 00 0 0 00000 00 00 0 000 00 000 0 0 00 0 0 00 00 0 0 000000 00 0000 0 0000 00 00 0 0 00 0 0 00 00 00 0 0 00 000 0 0 00 00000 0 00 00 0 0 0 00 000 0 00 00 0000 0 0000 00 00 0 00 00 0 0 0 00 000000 0 00 00 00 0 0 00 00 0 0 00 00000 0 00 00 0 0 0 00 0 0 0000 00 00 0 0 00 0 0 00000 00 00 0 0000 00 00 0 00 00 0 0 000 00 0 0 0 00 00 0 0 00 000000 0 00 00 00000 0 0 00 00000 0 00 00 0000 0 000 00 0 0 000 00 0 0 00 00 00 0 0 00 000 0 0 00 00000 0 000 00 0 0 00000 00 0 0 0 00 000 0 00 00 0 0 0 00 00 0 0000 00 0 0 0 00 00 0 00 00 00 0 0 00 0 0 0 00 0 0 0 00 00000 0 000 00 00 0 00000 00 0000 0 00 00 0000 0 000 00 000 0 0000 00 00 0 0 00 0 0 0 00 0 0 0 00 0 0 000 00 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChuckNorris(tt.message); got != tt.want {
				t.Errorf("ChuckNorris(%v) = %v, want %v", tt.message, got, tt.want)
			}
		})
	}
}

func Test_typeBlockFor(t *testing.T) {
	tests := []struct {
		name string
		bit  rune
		want string
	}{
		{"zero", '0', "00"},
		{"one", '1', "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeBlockFor(tt.bit); got != tt.want {
				t.Errorf("typeBlockFor(%q) = %v, want %v", tt.bit, got, tt.want)
			}
		})
	}
}
