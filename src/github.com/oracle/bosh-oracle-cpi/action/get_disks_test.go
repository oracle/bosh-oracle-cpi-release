package action

import (
	"errors"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/disks"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"

	clientfakes "github.com/oracle/bosh-oracle-cpi/bmc/client/fakes"
	diskfakes "github.com/oracle/bosh-oracle-cpi/bmc/disks/fakes"
)

var _ = Describe("GetDisks", func() {
	var (
		err             error
		volumes         []string
		attachedVolumes []*resource.Volume

		connector  *clientfakes.FakeConnector
		logger     boshlog.Logger
		diskFinder *diskfakes.FakeDiskFinder

		getDisks GetDisks
	)

	BeforeEach(func() {
		diskFinder = &diskfakes.FakeDiskFinder{}
		installDiskFinderFactory(func(c client.Connector, l boshlog.Logger, loc resource.Location) disks.Finder {
			return diskFinder
		})

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		getDisks = NewGetDisks(connector, logger)
	})

	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		Context("when there are attached disk", func() {

			BeforeEach(func() {
				attachedVolumes = []*resource.Volume{
					resource.NewVolume("fake-vol1", "fake-vol1-ocid"),
					resource.NewVolume("fake-vol2", "fake-vol2-ocid"),
				}
				diskFinder.FindAllAttachedResult = attachedVolumes
			})

			It("returns the list of attached disks", func() {
				volumes, err = getDisks.Run("fake-vm-ocid")
				Expect(err).NotTo(HaveOccurred())
				Expect(diskFinder.FindAllAttachedVolumesCalled).To(BeTrue())
				Expect(volumes).To(Equal([]string{"fake-vol1-ocid", "fake-vol2-ocid"}))
			})
		})

		Context("when there are not any attached disk", func() {
			It("returns an empty array", func() {
				volumes, err = getDisks.Run("fake-vm-ocid")
				Expect(err).NotTo(HaveOccurred())
				Expect(diskFinder.FindAllAttachedVolumesCalled).To(BeTrue())
				Expect(volumes).To(BeEmpty())
			})
		})

		It("returns an error if finder fails", func() {
			diskFinder.FindAllAttachedError = errors.New("fake-volfinder-error")

			_, err = getDisks.Run("fake-vm-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-volfinder-error"))
			Expect(diskFinder.FindAllAttachedVolumesCalled).To(BeTrue())
		})
	})
})
