package action

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	clientfakes "github.com/oracle/bosh-oracle-cpi/oci/client/fakes"
	vmfakes "github.com/oracle/bosh-oracle-cpi/oci/vm/fakes"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/vm"
)

var _ = Describe("SetVMMetaData", func() {
	var (
		connector *clientfakes.FakeConnector
		logger    boshlog.Logger
		updater   *vmfakes.FakeVMUpdater

		err error

		setVMMetadata SetVMMetadata
	)

	BeforeEach(func() {

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		updater = &vmfakes.FakeVMUpdater{}
		installVMUpdaterFactory(func(c client.Connector, l boshlog.Logger) vm.Updater {
			return updater
		})

		setVMMetadata = NewSetVMMetadata(connector, logger)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		It("calls updater if name exists in metadata", func() {

			md := VMMetadata{
				"name": "new-name",
			}
			_, err = setVMMetadata.Run("fake-ocid", md)
			Expect(updater.UpdateInstanceCalled).To(BeTrue())
			Expect(err).NotTo(HaveOccurred())
			Expect(updater.UpdatedName).To(Equal("new-name"))
		})

		It("doesn't fail or call updater if name doesn't exist in metadata", func() {
			_, err = setVMMetadata.Run("fake-ocid", VMMetadata{})
			Expect(updater.UpdateInstanceCalled).To(BeFalse())
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns error if updater returns an error", func() {
			updater.UpdateInstanceError = errors.New("fake-updater-error")

			_, err = setVMMetadata.Run("fake-ocid", VMMetadata{
				"name": "new-name",
			})
			Expect(updater.UpdateInstanceCalled).To(BeTrue())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-updater-error"))
		})
	})
})
