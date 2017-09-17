package disks

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDisks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Disks Suite")
}
