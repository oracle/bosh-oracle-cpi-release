package action

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/disks"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"github.com/oracle/bosh-oracle-cpi/bmc/stemcell"
	"github.com/oracle/bosh-oracle-cpi/bmc/vm"
)

var (
	stemcellCreatorFactory   stemcell.CreatorFactory   = stemcell.NewCreator
	stemcellDestroyerFactory stemcell.DestroyerFactory = stemcell.NewDestroyer

	vmFinderFactory     vm.FinderFactory     = vm.NewFinder
	vmCreatorFactory    vm.CreatorFactory    = vm.NewCreator
	vmTerminatorFactory vm.TerminatorFactory = vm.NewTerminator

	diskCreatorFactory      disks.CreatorFactory                  = disks.NewCreator
	diskFinderFactory       disks.FinderFactory                   = disks.NewFinder
	diskTerminatorFactory   disks.TerminatorFactory               = disks.NewTerminator
	attacherDetacherFactory disks.InstanceAttacherDetacherFactory = disks.NewAttacherDetacherForInstance
)

func newStemcellCreator(c client.Connector, l boshlog.Logger) stemcell.Creator {
	return stemcellCreatorFactory(c, l)
}

func newStemcellDestroyer(c client.Connector, l boshlog.Logger) stemcell.Destroyer {
	return stemcellDestroyerFactory(c, l)
}

func newVMFinder(c client.Connector, l boshlog.Logger) vm.Finder {
	return vmFinderFactory(c, l)
}

func newVMCreator(c client.Connector, l boshlog.Logger, vcnName string, subnetName string, availabilityDomain string) vm.Creator {
	return vmCreatorFactory(c, l, vcnName, subnetName, availabilityDomain)
}

func newVMTerminator(c client.Connector, l boshlog.Logger) vm.Terminator {
	return vmTerminatorFactory(c, l)
}

func newDiskCreator(c client.Connector, l boshlog.Logger, loc resource.Location) disks.Creator {
	return diskCreatorFactory(c, l, loc)
}

func newDiskFinder(c client.Connector, l boshlog.Logger, loc resource.Location) disks.Finder {
	return diskFinderFactory(c, l, loc)
}

func newDiskTerminator(c client.Connector, l boshlog.Logger) disks.Terminator {
	return diskTerminatorFactory(c, l)
}

func newAttacherDetacherForInstance(in *resource.Instance, c client.Connector, l boshlog.Logger) (disks.AttacherDetacher, error) {
	return attacherDetacherFactory(in, c, l)
}

func installStemcellCreatorFactory(fac stemcell.CreatorFactory) {
	stemcellCreatorFactory = fac
}
func installStemcellDestroyerFactory(fac stemcell.DestroyerFactory) {
	stemcellDestroyerFactory = fac
}

func installVMCreatorFactory(fac vm.CreatorFactory) {
	vmCreatorFactory = fac
}
func installVMTerminatorFactory(fac vm.TerminatorFactory) {
	vmTerminatorFactory = fac
}
func installVMFinderFactory(fac vm.FinderFactory) {
	vmFinderFactory = fac
}

func installDiskCreatorFactory(fac disks.CreatorFactory) {
	diskCreatorFactory = fac
}
func installDiskFinderFactory(fac disks.FinderFactory) {
	diskFinderFactory = fac
}
func installDiskTerminatorFactory(fac disks.TerminatorFactory) {
	diskTerminatorFactory = fac
}

func installInstanceAttacherDetacherFactory(fac disks.InstanceAttacherDetacherFactory) {
	attacherDetacherFactory = fac
}

func resetAllFactories() {
	stemcellCreatorFactory = stemcell.NewCreator
	stemcellDestroyerFactory = stemcell.NewDestroyer

	vmFinderFactory = vm.NewFinder
	vmCreatorFactory = vm.NewCreator
	vmTerminatorFactory = vm.NewTerminator

	diskCreatorFactory = disks.NewCreator
	diskFinderFactory = disks.NewFinder
	diskTerminatorFactory = disks.NewTerminator
	attacherDetacherFactory = disks.NewAttacherDetacherForInstance

}
