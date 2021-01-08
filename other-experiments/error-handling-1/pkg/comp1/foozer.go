package comp1

import (
	"fmt"

	"github.com/mig4/gogiggle/other-experiments/error-handling-1/pkg/comp2"
)

func FoozIt() ([]string, error) {
	// let's say it we arrive at the list of items to process internally here
	items := []string{"foo-ok1", "foo-ok2", "foo-ko", "foo-ok3"}

	var rItems []string
	for _, itm := range items {
		rItm, err := comp2.BoozIt(itm)
		if err != nil {
			return rItems, fmt.Errorf("error processing \"%s\": %w", itm, err)
		}
		rItems = append(rItems, rItm)
	}

	return rItems, nil
}
