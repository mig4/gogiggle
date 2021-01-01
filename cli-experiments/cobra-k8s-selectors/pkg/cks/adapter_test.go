package cks_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/cks"
	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"
)

// Names of characters expected in the in-memory repository, used in assertions.
const (
	Luke   = "Luke Skywalker"
	C3PO   = "C-3PO"
	R2D2   = "R2-D2"
	Vader  = "Darth Vader"
	Leia   = "Leia Organa"
	ObiWan = "Obi-Wan Kenobi"
	Chewie = "Chewbacca"
	Han    = "Han Solo"
	Yoda   = "Yoda"
)

func getName(c swapi.Character) string {
	return c.Name
}

// CharacterNamed returns a matcher that expects the character to have `Name`
// attribute equal to specified value.
func CharacterNamed(name string) types.GomegaMatcher {
	return WithTransform(getName, Equal(name))
}

// ConsistOfCharacters returns a matcher that expects the value to `ConsistOf`
// only specified named characters.
func ConsistOfCharacters(names ...string) types.GomegaMatcher {
	var matchers []interface{}
	for _, name := range names {
		matchers = append(matchers, CharacterNamed(name))
	}
	return ConsistOf(matchers...)
}

// failingRepo is an implementation of swapi.Port that returns failures. For
// use in failure-mode tests.
type failingRepo struct{}

func (r *failingRepo) GetAll() ([]swapi.Character, error) {
	return []swapi.Character{}, errors.New("failed to get characters")
}

var _ = Describe("Adapter", func() {
	var (
		adapter cks.Port
		repo    swapi.Port
	)

	Describe("List", func() {
		BeforeEach(func() {
			// We can just use the in-memory repo for tests
			repo = swapi.NewInMemoryRepository()
		})

		JustBeforeEach(func() {
			adapter = cks.NewCks(repo)
		})

		It("Returns all objects if passed no options", func() {
			Expect(adapter.List(&cks.ListOptions{})).To(SatisfyAll(
				HaveLen(9),
				ConsistOfCharacters(Luke, C3PO, R2D2, Vader, Leia, ObiWan, Chewie, Han, Yoda),
			))
		})

		DescribeTable(
			"filters elements correctly",
			func(opts *cks.ListOptions, len int, names ...string) {
				Expect(adapter.List(opts)).To(SatisfyAll(
					HaveLen(len),
					ConsistOfCharacters(names...),
				))
			},
			Entry("returns everything without filters", &cks.ListOptions{}, 9, Luke, C3PO, R2D2, Vader, Leia, ObiWan, Chewie, Han, Yoda),
			Entry("returns results filtered by name label", &cks.ListOptions{LabelSelector: "name=C-3PO"}, 1, C3PO),
			Entry("returns results filtered by name field", &cks.ListOptions{FieldSelector: "name=R2-D2"}, 1, R2D2),
			Entry("returns results filtered by gender field", &cks.ListOptions{FieldSelector: "gender=n/a"}, 2, C3PO, R2D2),
			Entry("returns results filtered by height label expression", &cks.ListOptions{LabelSelector: "height>180"}, 3, Vader, ObiWan, Chewie),
			Entry(
				"returns results filtered by gender field and mass label expression",
				&cks.ListOptions{LabelSelector: "mass<80", FieldSelector: "gender=male"},
				3, Luke, ObiWan, Yoda,
			),
			Entry("returns results filtered by bool field = true", &cks.ListOptions{FieldSelector: "ghost=true"}, 5, Luke, Vader, ObiWan, Han, Yoda),
			Entry("returns results filtered by bool field = false", &cks.ListOptions{FieldSelector: "ghost=false"}, 4, C3PO, R2D2, Leia, Chewie),
			Entry("returns results filtered by bool label = true", &cks.ListOptions{LabelSelector: "ghost"}, 5, Luke, Vader, ObiWan, Han, Yoda),
			Entry("returns results filtered by bool field = false", &cks.ListOptions{LabelSelector: "!ghost"}, 4, C3PO, R2D2, Leia, Chewie),
			Entry("returns results filtered by presence of jedi tag", &cks.ListOptions{LabelSelector: "jedi"}, 3, Luke, ObiWan, Yoda),
			Entry("returns results filtered by presence of smuggler tag", &cks.ListOptions{LabelSelector: "smuggler"}, 2, Chewie, Han),
			Entry("returns results filtered by absence of a human tag", &cks.ListOptions{LabelSelector: "!human"}, 4, C3PO, R2D2, Chewie, Yoda),
			Entry(
				"returns nothing when label doesn't match",
				&cks.ListOptions{LabelSelector: "jedi,black-lightsaber", FieldSelector: "gender=male"},
				0,
			),
			Entry(
				"returns nothing when field selector doesn't match",
				&cks.ListOptions{LabelSelector: "forceSensitive", FieldSelector: "gender=male,hairColor=green"},
				0,
			),
		)

		It("fails when an invalid label selector is given", func() {
			_, err := adapter.List(&cks.ListOptions{LabelSelector: "!x=a"})
			Expect(err).ToNot(Succeed())
		})

		It("fails when an invalid field selector is given", func() {
			_, err := adapter.List(&cks.ListOptions{FieldSelector: "x"})
			Expect(err).ToNot(Succeed())
		})

		Context("when repo fails", func() {
			BeforeEach(func() {
				repo = &failingRepo{}
			})

			It("propagates the error", func() {
				_, err := adapter.List(&cks.ListOptions{})
				Expect(err).To(MatchError("failed to get characters"))
			})
		})
	})
})
