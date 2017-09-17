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
	"github.com/oracle/bosh-oracle-cpi/registry"

	clientfakes "github.com/oracle/bosh-oracle-cpi/oci/client/fakes"
	diskfakes "github.com/oracle/bosh-oracle-cpi/oci/disks/fakes"
	vmfakes "github.com/oracle/bosh-oracle-cpi/oci/vm/fakes"
	registryfakes "github.com/oracle/bosh-oracle-cpi/registry/fakes"
)

var _ = Describe("DetachDisk", func() {
	var (
		err        error
		detacherVM *resource.Instance

		vmLocation    resource.Location
		foundInstance *resource.Instance
		foundVolume   *resource.Volume

		expectedAgentSettings registry.AgentSettings

		registryClient *registryfakes.FakeClient
		connector      *clientfakes.FakeConnector
		logger         boshlog.Logger

		vmFinder *vmfakes.FakeVMFinder

		diskFinder       *diskfakes.FakeDiskFinder
		attacherDetacher *diskfakes.FakeAttacherDetacher

		detachDisk DetachDisk
	)
	BeforeEach(func() {
		vmFinder = &vmfakes.FakeVMFinder{}
		installVMFinderFactory(func(c client.Connector, l boshlog.Logger) vm.Finder {
			return vmFinder
		})

		diskFinder = &diskfakes.FakeDiskFinder{}
		installDiskFinderFactory(func(c client.Connector, l boshlog.Logger, loc resource.Location) disks.Finder {
			return diskFinder
		})

		attacherDetacher = &diskfakes.FakeAttacherDetacher{}
		installInstanceAttacherDetacherFactory(func(in *resource.Instance, c client.Connector, l boshlog.Logger) (disks.AttacherDetacher, error) {
			detacherVM = in
			return attacherDetacher, nil
		})

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)
		registryClient = &registryfakes.FakeClient{}

		detachDisk = NewDetachDisk(connector, logger, registryClient)

		vmLocation = resource.NewLocation("fake-vcn", "fake-subnet", "fake-ad1", "fake-compartment-id")
		foundInstance = resource.NewInstance("fake-vm-ocid", vmLocation)
		vmFinder.FindInstanceResult = foundInstance

		foundVolume = resource.NewVolume("fake-vol-name", "fake-vol-ocid")
		foundVolume.SetAttachmentID("fake-attachment-ocid")
		foundVolume.SetAttachmentIP("fake-attachement-IP")
		foundVolume.SetAttachmentIQN("fake-iqn")
		foundVolume.SetDevicePath("/dev/fake-path")
		diskFinder.FindVolumeResult = foundVolume

		registryClient.FetchSettings = registry.AgentSettings{
			Disks: registry.DisksSettings{
				Persistent: map[string]registry.PersistentSettings{
					"fake-vol-ocid": {
						ID:   "fake-vol-ocid",
						Path: "/dev/fake-path",
					},
				},
			},
		}
		expectedAgentSettings = registry.AgentSettings{
			Disks: registry.DisksSettings{
				Persistent: map[string]registry.PersistentSettings{},
			},
		}
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		It("finds the vm", func() {
			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(vmFinder.FindInstanceCalled).To(BeTrue())
			Expect(vmFinder.FindInstanceID).To(Equal("fake-vm-ocid"))
		})

		It("creates a detacher for the found vm", func() {
			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(detacherVM).To(Equal(vmFinder.FindInstanceResult))
		})

		It("finds the disk", func() {
			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
			Expect(diskFinder.FindVolumeID).To(Equal("fake-vol-ocid"))
		})

		It("detaches the disk", func() {
			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(attacherDetacher.DetachVolumeCalled).To(BeTrue())
			Expect(attacherDetacher.DetachedVolume).To(Equal(foundVolume))
		})

		It("udates the registry", func() {
			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(registryClient.UpdateCalled).To(BeTrue())
			Expect(registryClient.UpdateSettings).To(Equal(expectedAgentSettings))

		})
		It("returns an error if vmfinder fails", func() {
			vmFinder.FindInstanceError = errors.New("fake-instance-finder-error")

			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).To(HaveOccurred())
			Expect(vmFinder.FindInstanceCalled).To(BeTrue())
			Expect(err.Error()).To(ContainSubstring("fake-instance-finder-error"))
			Expect(attacherDetacher.DetachVolumeCalled).To(BeFalse())
			Expect(registryClient.FetchCalled).To(BeFalse())
			Expect(registryClient.UpdateCalled).To(BeFalse())
		})

		It("returns an error if diskfinder fails", func() {
			diskFinder.FindVolumeError = errors.New("fake-disk-finder-error")

			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-disk-finder-error"))
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
			Expect(attacherDetacher.DetachVolumeCalled).To(BeFalse())
			Expect(registryClient.FetchCalled).To(BeFalse())
			Expect(registryClient.UpdateCalled).To(BeFalse())
		})

		It("returns an error if detacher fails", func() {
			attacherDetacher.DetachmentError = errors.New("fake-attachment-error")

			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-attachment-error"))
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
			Expect(vmFinder.FindInstanceCalled).To(BeTrue())
			Expect(attacherDetacher.DetachVolumeCalled).To(BeTrue())
			Expect(registryClient.FetchCalled).To(BeFalse())
			Expect(registryClient.UpdateCalled).To(BeFalse())
		})

		It("returns an error if registryClient fetch call returns an error", func() {
			registryClient.FetchErr = errors.New("fake-registry-client-error")

			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-registry-client-error"))
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
			Expect(vmFinder.FindInstanceCalled).To(BeTrue())
			Expect(attacherDetacher.DetachVolumeCalled).To(BeTrue())
			Expect(registryClient.FetchCalled).To(BeTrue())
			Expect(registryClient.UpdateCalled).To(BeFalse())
		})

		It("returns an error if registryClient update call returns an error", func() {
			registryClient.UpdateErr = errors.New("fake-registry-client-error")

			_, err = detachDisk.Run("fake-vm-ocid", "fake-vol-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-registry-client-error"))
			Expect(diskFinder.FindVolumeCalled).To(BeTrue())
			Expect(vmFinder.FindInstanceCalled).To(BeTrue())
			Expect(attacherDetacher.DetachVolumeCalled).To(BeTrue())
			Expect(registryClient.FetchCalled).To(BeTrue())
			Expect(registryClient.UpdateCalled).To(BeTrue())
		})
	})
})
