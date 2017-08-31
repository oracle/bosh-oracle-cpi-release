package action

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/disks"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"github.com/oracle/bosh-oracle-cpi/bmc/vm"

	"errors"
	clientfakes "github.com/oracle/bosh-oracle-cpi/bmc/client/fakes"
	diskfakes "github.com/oracle/bosh-oracle-cpi/bmc/disks/fakes"
	vmfakes "github.com/oracle/bosh-oracle-cpi/bmc/vm/fakes"
	registryfakes "github.com/oracle/bosh-oracle-cpi/registry/fakes"
)

var _ = Describe("DeleteVM", func() {
	var (
		err            error
		registryClient *registryfakes.FakeClient
		connector      *clientfakes.FakeConnector
		logger         boshlog.Logger
		terminator     *vmfakes.FakeVMTerminator
		finder         *vmfakes.FakeVMFinder

		diskFinder       *diskfakes.FakeDiskFinder
		attacherDetacher *diskfakes.FakeAttacherDetacher

		deleteVM DeleteVM
	)

	BeforeEach(func() {

		terminator = &vmfakes.FakeVMTerminator{}
		installVMTerminatorFactory(func(c client.Connector, l boshlog.Logger) vm.Terminator {
			return terminator
		})

		finder = &vmfakes.FakeVMFinder{}
		installVMFinderFactory(func(c client.Connector, l boshlog.Logger) vm.Finder {
			return finder
		})

		diskFinder = &diskfakes.FakeDiskFinder{}
		installDiskFinderFactory(func(c client.Connector, l boshlog.Logger, loc resource.Location) disks.Finder {
			return diskFinder
		})

		attacherDetacher = &diskfakes.FakeAttacherDetacher{}
		installInstanceAttacherDetacherFactory(func(in *resource.Instance, c client.Connector, l boshlog.Logger) (disks.AttacherDetacher, error) {
			return attacherDetacher, nil
		})

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)
		registryClient = &registryfakes.FakeClient{}

		deleteVM = NewDeleteVM(connector, logger, registryClient)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		It("detaches any attached volumes, deletes the given instance, and updates the registry", func() {
			_, err = deleteVM.Run("fake-ocid")
			Expect(err).NotTo(HaveOccurred())
			Expect(diskFinder.FindAllAttachedVolumesCalled).To(BeTrue())
			Expect(terminator.TerminateInstanceCalled).To(BeTrue())
			Expect(terminator.TerminatedInstance).To(Equal("fake-ocid"))
			Expect(registryClient.DeleteCalled).To(BeTrue())

		})
		It("returns an error if instance terminator fails", func() {
			terminator.TerminateInstanceError = errors.New("fake-delete-error")

			_, err = deleteVM.Run("fake-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-delete-error"))
			Expect(terminator.TerminateInstanceCalled).To(BeTrue())
			Expect(registryClient.DeleteCalled).To(BeFalse())
		})
		It("returns an error if registryClient delete call returns an error", func() {
			registryClient.DeleteErr = errors.New("fake-registry-client-error")

			_, err = deleteVM.Run("fake-ocid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-registry-client-error"))
			Expect(diskFinder.FindAllAttachedVolumesCalled).To(BeTrue())
			Expect(terminator.TerminateInstanceCalled).To(BeTrue())
			Expect(registryClient.DeleteCalled).To(BeTrue())
		})
		Context("When disk findattachment or detachment fails", func() {
			It("Ignores those errors", func() {

				diskFinder.FindAllAttachedError = errors.New("fake-findallattachments-error")

				_, err = deleteVM.Run("fake-ocid")
				Expect(err).ToNot(HaveOccurred())
				Expect(diskFinder.FindAllAttachedVolumesCalled).To(BeTrue())
				Expect(terminator.TerminateInstanceCalled).To(BeTrue())
				Expect(registryClient.DeleteCalled).To(BeTrue())
			})

		})

	})
})
