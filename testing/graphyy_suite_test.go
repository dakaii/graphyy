package testing

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGraphyy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphyy Suite")
}
