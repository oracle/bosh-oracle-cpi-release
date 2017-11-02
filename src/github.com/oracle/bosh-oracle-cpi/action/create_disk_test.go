package action

import (
	"errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/disks"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"github.com/oracle/bosh-oracle-cpi/oci/vm"

	clientfakes "github.com/oracle/bosh-oracle-cpi/oci/client/fakes"
	diskfakes "github.com/oracle/bosh-oracle-cpi/oci/disks/fakes"
	vmfakes "github.com/oracle/bosh-oracle-cpi/oci/vm/fakes"
)

var _ = Describe("CreateDisk", func() {
	var (
		err        error
		diskCID    DiskCID
		vmCID      VMCID
		cloudProps DiskCloudProperties
		vmLocation resource.Location

		connector   *clientfakes.FakeConnector
		logger      boshlog.Logger
		diskCreator *diskfakes.FakeDiskCreator
		vmFinder    *vmfakes.FakeVMFinder

		createDisk CreateDisk
	)

	BeforeEach(func() {

		vmFinder = &vmfakes.FakeVMFinder{}
		installVMFinderFactory(func(c client.Connector, l boshlog.Logger) vm.Finder {
			return vmFinder
		})

		diskCreator = &diskfakes.FakeDiskCreator{}
		installDiskCreatorFactory(func(c client.Connector, l boshlog.Logger, loc resource.Location) disks.Creator {
			diskCreator.CreateVolumeLocation = loc
			return diskCreator
		})

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		createDisk = NewCreateDisk(connector, logger)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		Context("when vmCID instance is found", func() {

			BeforeEach(func() {
				vmLocation = resource.NewLocation("fake-vcn", "fake-subnet", "fake-ad1", "")
				vmFinder.FindInstanceResult = resource.NewInstance("fake-vm-ocid", vmLocation)
				cloudProps = DiskCloudProperties{}
			})

			It("creates the disk", func() {

				diskCreator.CreateVolumeResult = resource.NewVolume("fake-volume-name", "fake-volume-ocid")
				diskCID, err = createDisk.Run(51200, cloudProps, "fake-vm-ocid")

				Expect(err).NotTo(HaveOccurred())
				Expect(vmFinder.FindInstanceCalled).To(BeTrue())
				Expect(diskCreator.CreateVolumeCalled).To(BeTrue())
				Expect(diskCID).To(Equal(DiskCID("fake-volume-ocid")))
			})
			It("creates the disk in instance's avaialibilty domain", func() {

				diskCreator.CreateVolumeResult = resource.NewVolume("fake-volume-name", "fake-volume-ocid")
				diskCID, err = createDisk.Run(51200, cloudProps, "fake-vm-ocid")

				Expect(err).NotTo(HaveOccurred())
				Expect(vmFinder.FindInstanceCalled).To(BeTrue())
				Expect(diskCreator.CreateVolumeLocation.AvailabilityDomain()).To(Equal("fake-ad1"))
			})

			It("returns an error if disk creator fails", func() {
				diskCreator.CreateVolumeError = errors.New("fake-create-volume-error")

				_, err = createDisk.Run(51200, cloudProps, vmCID)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("fake-create-volume-error"))
				Expect(vmFinder.FindInstanceCalled).To(BeTrue())
			})

			It("rounds up the requested disk size to the minimum supported size", func() {
				diskCreator.CreateVolumeResult = resource.NewVolume("fake-volume-name", "fake-volume-ocid")

				_, err = createDisk.Run(2048, cloudProps, vmCID)
				Expect(err).NotTo(HaveOccurred())
				Expect(diskCreator.CreateVolumeCalled).To(BeTrue())
				Expect(diskCreator.CreateVolumeSize >= minVolumeSize).To(BeTrue())
			})

		})

		Context("when the given instance is not found", func() {
			BeforeEach(func() {
				vmCID = VMCID("fake-vm-cid")
				vmFinder.FindInstanceResult = nil
				vmFinder.FindInstanceError = errors.New("fake-findinstance-error")
			})

			It("bubbles up the error from vm finder", func() {
				diskCID, err = createDisk.Run(51200, cloudProps, vmCID)
				Expect(err).To(HaveOccurred())
				Expect(vmFinder.FindInstanceCalled).To(BeTrue())
				Expect(err.Error()).To(ContainSubstring("fake-findinstance-error"))
			})

		})
	})
})

var _ = Describe("volumeSize", func() {
	It("rounds up volume size to minimum 50GiB in 1GiB increments", func() {
		sizes := []struct {
			in  int
			out int64
		}{
			{0, 51200},
			{1024, 51200},
			{2048, 51200},
			{51200, 51200},
			{51201, 52224},
			{51200 + 2047, 51200 + 2048},
			{51200 * 2, 51200 * 2},
		}
		for _, s := range sizes {
			Expect(volumeSize(s.in)).To(Equal(s.out))
		}
	})
})
