package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type=WorldStatus

type WorldStatus int

const (
	Unknown WorldStatus = iota
	Burning
	Calm
	Beautiful
)

func main() {
	fmt.Printf("Hello %v World\n", Beautiful)
}
