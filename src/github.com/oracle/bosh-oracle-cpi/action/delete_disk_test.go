package action

import (
	"errors"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	clientfakes "github.com/oracle/bosh-oracle-cpi/oci/client/fakes"
	"github.com/oracle/bosh-oracle-cpi/oci/disks"
	diskfakes "github.com/oracle/bosh-oracle-cpi/oci/disks/fakes"
)

var _ = Describe("DeleteDisk", func() {
	var (
		err error

		connector      *clientfakes.FakeConnector
		logger         boshlog.Logger
		diskTerminator *diskfakes.FakeDiskTerminator

		deleteDisk DeleteDisk
	)

	BeforeEach(func() {
		diskTerminator = &diskfakes.FakeDiskTerminator{}
		installDiskTerminatorFactory(func(c client.Connector, l boshlog.Logger) disks.Terminator {
			return diskTerminator
		})

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		deleteDisk = NewDeleteDisk(connector, logger)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		It("delegates to disk terminator", func() {
			_, err = deleteDisk.Run("fake-volume-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(diskTerminator.DeleteVolumeCalled).To(BeTrue())
			Expect(diskTerminator.DeleteVolumeID).To(Equal("fake-volume-ocid"))
		})

		It("returns an error if disk terminator returns an error", func() {
			diskTerminator.DeleteVolumeError = errors.New("fake-delete-volume-error")

			_, err = deleteDisk.Run("fake-volume-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-delete-volume-error"))
			Expect(diskTerminator.DeleteVolumeCalled).To(BeTrue())
		})
	})
})
