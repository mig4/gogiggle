package cks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/fields"

	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/cks"
	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"
)

var _ = Describe("Fields", func() {
	Describe("CharacterAsFields", func() {
		It("exposes Character fields as a Set of fields", func() {
			Expect(cks.CharacterAsFields(swapi.Character{
				Name:           "Foo Bar",
				Height:         150,
				Mass:           80,
				HairColor:      "blue",
				Gender:         "male",
				ForceSensitive: true,
				Ghost:          true,
			})).To(Equal(fields.Set{
				"name":           "Foo Bar",
				"height":         "150",
				"mass":           "80",
				"hairColor":      "blue",
				"gender":         "male",
				"forceSensitive": "true",
				"ghost":          "true",
				"tags":           "",
			}))
		})

		It("treats false bool fields same as other fields", func() {
			Expect(cks.CharacterAsFields(swapi.Character{
				Name:           "Foo Baz",
				Height:         170,
				Mass:           100,
				HairColor:      "red",
				Gender:         "female",
				ForceSensitive: false,
				Ghost:          false,
			})).To(Equal(fields.Set{
				"name":           "Foo Baz",
				"height":         "170",
				"mass":           "100",
				"hairColor":      "red",
				"gender":         "female",
				"forceSensitive": "false",
				"ghost":          "false",
				"tags":           "",
			}))
		})

		It("exposes unset fields as zero-values", func() {
			Expect(cks.CharacterAsFields(swapi.Character{
				Name: "Foo Qux",
			})).To(Equal(fields.Set{
				"name":           "Foo Qux",
				"height":         "0",
				"mass":           "0",
				"hairColor":      "",
				"gender":         "",
				"forceSensitive": "false",
				"ghost":          "false",
				"tags":           "",
			}))
		})

		It("exposes tags as a field with comma-separated values", func() {
			Expect(cks.CharacterAsFields(swapi.Character{
				Name:      "Foo Quux",
				Height:    150,
				Mass:      80,
				HairColor: "blue",
				Gender:    "male",
				Tags:      []string{"tag1", "tag2", "gender"},
			})).To(Equal(fields.Set{
				"name":           "Foo Quux",
				"height":         "150",
				"mass":           "80",
				"hairColor":      "blue",
				"gender":         "male",
				"forceSensitive": "false",
				"ghost":          "false",
				"tags":           "tag1,tag2,gender",
			}))
		})
	})
})
