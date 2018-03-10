package action

import (
	"errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	clientfakes "github.com/oracle/bosh-oracle-cpi/oci/client/fakes"
	stemcellfakes "github.com/oracle/bosh-oracle-cpi/oci/stemcell/fakes"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/stemcell"
)

var _ = Describe("DeleteStemcell", func() {
	var (
		err error

		connector *clientfakes.FakeConnector
		logger    boshlog.Logger
		destroyer *stemcellfakes.FakeDestroyer

		deleteStemcell DeleteStemcell
	)

	BeforeEach(func() {
		destroyer = &stemcellfakes.FakeDestroyer{}
		installStemcellDestroyerFactory(func(c client.Connector, l boshlog.Logger) stemcell.Destroyer {
			return destroyer
		})

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		deleteStemcell = NewDeleteStemcell(connector, logger)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		It("delegates to StemCell Destroyer", func() {
			_, err = deleteStemcell.Run("fake-stemcell-id")
			Expect(err).NotTo(HaveOccurred())
			Expect(destroyer.DestroyStemcellCalled).To(BeTrue())
		})

		It("returns error if Stemcell Destroyer fails", func() {
			destroyer.DestroyStemcellError = errors.New("fake-stemcell-error")

			_, err = deleteStemcell.Run("fake-stemcell-id")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-stemcell-error"))
			Expect(destroyer.DestroyStemcellCalled).To(BeTrue())
		})

	})
})
