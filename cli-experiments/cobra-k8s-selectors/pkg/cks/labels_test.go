package cks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/cks"
	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"
)

var _ = Describe("Labels", func() {
	Describe("CharacterAsLabels", func() {
		It("exposes Character fields as a Set of labels", func() {
			Expect(cks.CharacterAsLabels(swapi.Character{
				Name:           "Foo Bar",
				Height:         150,
				Mass:           80,
				HairColor:      "blue",
				Gender:         "male",
				ForceSensitive: true,
				Ghost:          true,
			})).To(Equal(labels.Set{
				"name":           "Foo Bar",
				"height":         "150",
				"mass":           "80",
				"hairColor":      "blue",
				"gender":         "male",
				"forceSensitive": "true",
				"ghost":          "true",
			}))
		})

		It("treats false bool fields as absent labels", func() {
			Expect(cks.CharacterAsLabels(swapi.Character{
				Name:           "Foo Baz",
				Height:         170,
				Mass:           100,
				HairColor:      "red",
				Gender:         "female",
				ForceSensitive: false,
				Ghost:          false,
			})).To(Equal(labels.Set{
				"name":      "Foo Baz",
				"height":    "170",
				"mass":      "100",
				"hairColor": "red",
				"gender":    "female",
			}))
		})

		It("exposes unset fields as zero-values", func() {
			Expect(cks.CharacterAsLabels(swapi.Character{
				Name: "Foo Qux",
			})).To(Equal(labels.Set{
				"name":      "Foo Qux",
				"height":    "0",
				"mass":      "0",
				"hairColor": "",
				"gender":    "",
			}))
		})

		It("exposes tags as bool enabled labels but avoids overriding fields", func() {
			Expect(cks.CharacterAsLabels(swapi.Character{
				Name:      "Foo Quux",
				Height:    150,
				Mass:      80,
				HairColor: "blue",
				Gender:    "male",
				Tags:      []string{"tag1", "tag2", "gender"},
			})).To(Equal(labels.Set{
				"name":      "Foo Quux",
				"height":    "150",
				"mass":      "80",
				"hairColor": "blue",
				"gender":    "male",
				"tag1":      "true",
				"tag2":      "true",
			}))
		})
	})
})
