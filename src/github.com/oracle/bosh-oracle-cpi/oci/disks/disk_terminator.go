package disks

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/baremetal/core/client/blockstorage"
)

type diskTerminator struct {
	connector client.Connector
	logger    boshlog.Logger
}

func NewTerminator(c client.Connector, l boshlog.Logger) Terminator {
	return &diskTerminator{connector: c, logger: l}
}

type TerminatorFactory func(client.Connector, boshlog.Logger) Terminator

func (dt *diskTerminator) DeleteVolume(volumeID string) error {

	p := blockstorage.NewDeleteVolumeParams().WithVolumeID(volumeID)

	_, err := dt.connector.CoreSevice().Blockstorage.DeleteVolume(p)
	if err != nil {
		dt.logger.Error(diskOperationsLogTag, "Error deleting volume %v", err)
		return err
	}
	dt.logger.Debug(diskOperationsLogTag, "Deleted volume %s", volumeID)
	return nil
}
