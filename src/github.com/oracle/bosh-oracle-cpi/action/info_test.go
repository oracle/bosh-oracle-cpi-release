package action

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info", func() {
	var infoAction Info

	BeforeEach(func() {
		infoAction = NewInfo()
	})

	Describe("Run", func() {
		Context("stemcell_formats", func() {
			var info StemcellInfo

			BeforeEach(func() {
				var err error

				info, err = infoAction.Run()
				Expect(err).NotTo(HaveOccurred())
			})

			It("supports oracle-light", func() {
				Expect(info.Formats).To(ContainElement("oracle-light"))
			})
		})
	})
})
