package action

import (
	"errors"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/disks"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"

	clientfakes "github.com/oracle/bosh-oracle-cpi/oci/client/fakes"
	diskfakes "github.com/oracle/bosh-oracle-cpi/oci/disks/fakes"
)

var _ = Describe("HasDisk", func() {
	var (
		err   error
		found bool

		connector  *clientfakes.FakeConnector
		logger     boshlog.Logger
		diskFinder *diskfakes.FakeDiskFinder

		hasDisk HasDisk
	)

	BeforeEach(func() {
		diskFinder = &diskfakes.FakeDiskFinder{}
		installDiskFinderFactory(func(c client.Connector, l boshlog.Logger, loc resource.Location) disks.Finder {
			return diskFinder
		})

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		hasDisk = NewHasDisk(connector, logger)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		It("returns true if disk exists", func() {

			diskFinder.FindVolumeResult = resource.NewVolume("fake-name", "fake-volume-ocid")

			found, err = hasDisk.Run("fake-volume-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
			Expect(diskFinder.FindVolumeID).To(Equal("fake-volume-ocid"))
		})

		It("returns false if disk ID does not exist", func() {

			diskFinder.FindVolumeResult = resource.NewVolume("fake-name", "other-vol-ocid")

			found, err = hasDisk.Run("fake-volume-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeFalse())
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
		})

		It("returns an error if disk finder fails", func() {
			diskFinder.FindVolumeError = errors.New("fake-find-vol-error")

			_, err = hasDisk.Run("fake-volume-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-find-vol-error"))
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
		})
	})
})
