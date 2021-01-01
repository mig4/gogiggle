package cks

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/fields"

	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"
)

// CharacterAsFields transforms a Character struct into a field Set that
// implements the Fields interface.
func CharacterAsFields(c swapi.Character) fields.Fields {
	return fields.Set{
		"name":           c.Name,
		"height":         fmt.Sprint(c.Height),
		"mass":           fmt.Sprint(c.Mass),
		"hairColor":      c.HairColor,
		"gender":         c.Gender,
		"forceSensitive": fmt.Sprint(c.ForceSensitive),
		"ghost":          fmt.Sprint(c.Ghost),
		"tags":           strings.Join(c.Tags, ","),
	}
}
