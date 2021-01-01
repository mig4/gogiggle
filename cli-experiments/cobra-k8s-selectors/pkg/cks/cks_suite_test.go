package cks_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cks Suite")
}
