package comp2

import (
	"errors"
	"fmt"
	"strings"
)

func BoozIt(s string) (string, error) {
	if strings.HasSuffix(s, "ko") {
		return s, errors.New("got ko-d")
	}
	return fmt.Sprintf("boozing-%s", s), nil
}
