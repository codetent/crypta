package crypta_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var pathToCrypta string

func TestCrypta(t *testing.T) {
	BeforeSuite(func() {
		var err error
		pathToCrypta, err = gexec.Build("github.com/codetent/crypta")
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Crypta Suite")
}
