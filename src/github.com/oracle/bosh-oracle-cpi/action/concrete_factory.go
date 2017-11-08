package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"

	"github.com/oracle/bosh-oracle-cpi/config"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/registry"
)

// ConcreteFactory creates Actions
// to handle CPI requests from BOSH
type ConcreteFactory struct {
	connector        client.Connector
	availableActions map[string]Action
}

// NewConcreteFactory creates a ConcreteFactory instance
func NewConcreteFactory(
	connector client.Connector,
	uuidGen boshuuid.Generator,
	cfg config.Config,
	logger boshlog.Logger,
) ConcreteFactory {

	registryClient := registry.NewHTTPClient(cfg.Cloud.Properties.Registry, logger)

	return ConcreteFactory{
		connector: connector,
		availableActions: map[string]Action{
			// Image management
			"create_stemcell": NewCreateStemcell(connector, logger),
			"delete_stemcell": NewDeleteStemcell(connector, logger),

			// VM management
			"create_vm":       NewCreateVM(connector, logger, registryClient, uuidGen),
			"delete_vm":       NewDeleteVM(connector, logger, registryClient),
			"has_vm":          NewHasVM(connector, logger),
			"set_vm_metadata": NewSetVMMetadata(connector, logger),

			// Disk Management
			"create_disk": NewCreateDisk(connector, logger),
			"delete_disk": NewDeleteDisk(connector, logger),
			"attach_disk": NewAttachDisk(connector, logger, registryClient),
			"detach_disk": NewDetachDisk(connector, logger, registryClient),
			"has_disk":    NewHasDisk(connector, logger),
			"get_disks":   NewGetDisks(connector, logger),
		},
	}
}

// Create creates an Action responsible for handling
// a method request
func (f ConcreteFactory) Create(method string) (Action, error) {
	a, found := f.availableActions[method]
	if !found {
		return nil, bosherr.Errorf("Could not create action with method %s", method)
	}
	if err := f.connector.Connect(); err != nil {
		return nil, err
	}
	return a, nil
}
