package action

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
)

// DeleteStemcell action handles the delete_stemcell method invocation
type DeleteStemcell struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewDeleteStemcell creates a DeleteStemcell instance
func NewDeleteStemcell(connector client.Connector, logger boshlog.Logger) DeleteStemcell {
	return DeleteStemcell{connector, logger}
}

// Run delegates the request to a StemCellDestroyer
func (ds DeleteStemcell) Run(stemcellCID StemcellCID) (interface{}, error) {
	d := newStemcellDestroyer(ds.connector, ds.logger)
	if err := d.DeleteStemcell(string(stemcellCID)); err != nil {
		return nil, err
	}
	return nil, nil
}
