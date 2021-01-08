package main

import (
	"fmt"
	"os"

	"github.com/mig4/gogiggle/other-experiments/error-handling-1/pkg/comp1"
)

func main() {
	fmt.Println("calling comp1")
	items, err := comp1.FoozIt()
	for _, itm := range items {
		fmt.Printf("processed \"%s\"\n", itm)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "got err: %s\n", err)
	}
}
