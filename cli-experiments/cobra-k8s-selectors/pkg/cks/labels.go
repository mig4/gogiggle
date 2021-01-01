package cks

import (
	"fmt"

	"k8s.io/apimachinery/pkg/labels"

	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"
)

// CharacterAsLabels transforms a Character struct into a label Set that
// implements the Labels interface.
func CharacterAsLabels(c swapi.Character) labels.Labels {
	ls := labels.Set{
		"name":      c.Name,
		"height":    fmt.Sprint(c.Height),
		"mass":      fmt.Sprint(c.Mass),
		"hairColor": c.HairColor,
		"gender":    c.Gender,
	}

	if c.ForceSensitive {
		ls["forceSensitive"] = fmt.Sprint(c.ForceSensitive)
	}

	if c.Ghost {
		ls["ghost"] = fmt.Sprint(c.Ghost)
	}

	for _, tag := range c.Tags {
		if _, ok := ls[tag]; !ok {
			ls[tag] = "true"
		}
	}

	return ls
}
