package cks

import (
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"
)

type cks struct {
	repo swapi.Port
}

func NewCks(r swapi.Port) *cks {
	return &cks{repo: r}
}

// List returns a list of Characters from the repo, optionally filtered by
// a label expression and/or a field expression.
func (c *cks) List(options *ListOptions) ([]swapi.Character, error) {
	allCharacters, err := c.repo.GetAll()
	if err != nil {
		return []swapi.Character{}, err
	}

	ls := labels.Everything()
	if options.LabelSelector != "" {
		ls, err = labels.Parse(options.LabelSelector)
		if err != nil {
			return []swapi.Character{}, err
		}
	}

	fs := fields.Everything()
	if options.FieldSelector != "" {
		fs, err = fields.ParseSelector(options.FieldSelector)
		if err != nil {
			return []swapi.Character{}, err
		}
	}

	var characters []swapi.Character
	for _, c := range allCharacters {
		if !ls.Matches(CharacterAsLabels(c)) {
			continue
		}

		if !fs.Matches(CharacterAsFields(c)) {
			continue
		}

		characters = append(characters, c)
	}

	return characters, nil
}
