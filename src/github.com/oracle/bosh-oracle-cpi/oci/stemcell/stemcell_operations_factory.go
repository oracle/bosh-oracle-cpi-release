package stemcell

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
)

const stemCellLogTag = "OCIStemcell"

type Creator interface {
	CreateStemcell(imageOCID string) (stemcellId string, err error)
}
type CreatorFactory func(client.Connector, boshlog.Logger) Creator

type Destroyer interface {
	DeleteStemcell(stemcellId string) (err error)
}
type DestroyerFactory func(client.Connector, boshlog.Logger) Destroyer

func NewCreator(c client.Connector, l boshlog.Logger) Creator {
	return stemcellOperations{connector: c, logger: l}
}

func NewDestroyer(c client.Connector, l boshlog.Logger) Destroyer {
	return stemcellOperations{connector: c, logger: l}
}
