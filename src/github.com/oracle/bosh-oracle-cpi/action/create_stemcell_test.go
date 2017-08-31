package action

import (
	"errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	clientfakes "github.com/oracle/bosh-oracle-cpi/bmc/client/fakes"
	stemcellfakes "github.com/oracle/bosh-oracle-cpi/bmc/stemcell/fakes"

	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/stemcell"
)

var _ = Describe("CreateStemcell", func() {
	var (
		err        error
		cloudProps StemcellCloudProperties

		connector      *clientfakes.FakeConnector
		logger         boshlog.Logger
		creator        *stemcellfakes.FakeCreator
		createStemcell CreateStemcell
	)

	BeforeEach(func() {

		creator = &stemcellfakes.FakeCreator{}
		installStemcellCreatorFactory(func(c client.Connector, l boshlog.Logger) stemcell.Creator {
			return creator
		})
		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		createStemcell = NewCreateStemcell(connector, logger)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		BeforeEach(func() {
			cloudProps = StemcellCloudProperties{
				ImageOCID: "fake-image-ocid",
			}
		})
		It("delegates to StemCell Creator", func() {
			_, err = createStemcell.Run("", cloudProps)
			Expect(err).NotTo(HaveOccurred())
			Expect(creator.CreateStemcellCalled).To(BeTrue())
			Expect(creator.CreateStemcellCalledWithID).To(Equal("fake-image-ocid"))
		})

		It("returns error if stemcell creator fails ", func() {
			creator.CreateStemcellError = errors.New("fake-create-stemcell-error")

			_, err = createStemcell.Run("", cloudProps)
			Expect(err).To(HaveOccurred())
			Expect(creator.CreateStemcellCalled).To(BeTrue())
			Expect(err.Error()).To(ContainSubstring("fake-create-stemcell-error"))
		})
	})

})
