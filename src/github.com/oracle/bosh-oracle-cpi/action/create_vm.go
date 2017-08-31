package action

import (
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"

	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"time"
)

// CreateVM action handles the create_vm request
type CreateVM struct {
	connector client.Connector
	logger    boshlog.Logger
	registry  registry.Client
	uuidGen   boshuuid.Generator
}

const logTag = "createVM"

// NewCreateVM creates a CreateVM instance
func NewCreateVM(c client.Connector, l boshlog.Logger, r registry.Client, u boshuuid.Generator) CreateVM {
	return CreateVM{connector: c, logger: l, registry: r, uuidGen: u}
}

// Run creates an instance for the given configuration in BMC.
//
// If the instance is configured to be created in a manual network it assigns the
// given private IP to that instance.
//
// For dynamic network where the instance is not assigned an IP,
// it queries the IPs assigned to the new instance. In addition, if a SSHTunnel is configured
// it creates a forward tunnel to the public IP of that instance.
//
// Finally, it updates the agent registry with details of the new instance
func (cv CreateVM) Run(agentID string, stemcellCID StemcellCID, cloudProps VMCloudProperties,
	networks Networks, _ []DiskCID, env Environment) (VMCID, error) {

	// network properties
	n, err := cv.selectNetwork(networks)
	if err != nil {
		return "", err
	}
	agentNetworks := networks.AsRegistryNetworks()

	// Create the VM
	name := cv.vmName()
	creator := newVMCreator(cv.connector, cv.logger,
		n.CloudProperties.VcnName,
		n.CloudProperties.SubnetName, cloudProps.AvailabilityDomain)

	icfg := vm.InstanceConfiguration{
		Name:      name,
		Shape:     cloudProps.Shape,
		ImageId:   string(stemcellCID),
		PrivateIP: n.IP,
	}
	metadata := vm.InstanceMetadata{
		vm.NewSSHKeys(cv.connector.AuthorizedKeys()),
		vm.NewUserData(name, cv.connector.AgentRegistryEndpoint(),
			n.DNS, agentNetworks),
	}
	instance, err := creator.CreateInstance(icfg, metadata)

	// Optimization: Only do this if we must as in the case of dynamic
	// network configuration, since this starts ssh waits and pings to the agent.
	// However, the network type property is not always known
	// (bosh cli v2 gobbles it and doesn't pass it down to us).
	// So we guess -- dynamic configuration typically doesn't request a static private
	// IP and we use that as an indicator
	if err == nil && n.IP == "" {
		err = instance.EnsureReachable(cv.connector, cv.logger)
	}
	if err != nil {
		return "", bosherr.WrapError(err, "Error launching new instance")
	}

	if err := cv.updateRegistry(agentID, instance.ID(), name, agentNetworks, env); err != nil {
		return "", err
	}
	return VMCID(instance.ID()), nil
}

// selectNetwork finds a suitable network (to create the VM)
// from the list of networks configured in the
// deployment manifest
func (cv CreateVM) selectNetwork(networks Networks) (*Network, error) {

	// Prefer manual network over dynamic
	// network. Pick the first one
	n := networks.FirstStatic()
	if n == nil {
		n = networks.FirstDynamic()
	}
	// CF-51. Pick first network if type is missing
	// and we can't determine the network type
	if n == nil {
		n = networks.First()
	}
	// validate
	if n == nil {
		return nil, bosherr.Error("No suitable network defintion for creating a VM")
	}
	if err := n.validate(); err != nil {
		return n, err
	}
	return n, nil
}

func (cv CreateVM) updateRegistry(agentID string, instanceID string, vmName string,
	agentNetworks registry.NetworksSettings, env Environment) error {

	// Handle registry update failure. Delete VM or retry?
	var err error
	defer func() {
		if err != nil {
			cv.logger.Error(logTag, "Registry update failed! FIXME: handle failure")
		}
	}()
	agentOptions := cv.connector.AgentOptions()
	agentSettings := registry.NewAgentSettings(agentID, vmName, agentNetworks,
		registry.EnvSettings(env), agentOptions)

	// Update Registry with AgentSettings
	// for the agent (agent will find them as a HTTP source)
	if err = cv.registry.Update(instanceID, agentSettings); err != nil {
		return bosherr.WrapError(err, "Create VM. Error updating registry")
	}
	return nil

}

func (cv CreateVM) vmName() string {

	suffix, err := cv.uuidGen.Generate()
	if err != nil {
		suffix = time.Now().Format(time.Stamp)
	}
	return fmt.Sprintf("bosh-%s", suffix)
}
