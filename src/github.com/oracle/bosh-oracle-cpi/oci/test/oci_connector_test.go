package test

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/baremetal/core/client/compute"

	"testing"
)

// Tests Connector.Connect()
func Test_ConnectorConnect(t *testing.T) {

	cfg, ini, err := NewTestConfig(iniConfigPath(), iniTestSection())
	if err != nil {
		t.Fatalf("Error loading %s .%v", iniConfigPath(), err)
	}
	var ad = ini.AvailabilityDomain

	connector := client.NewConnector(cfg, boshlog.NewLogger(boshlog.LevelWarn))
	err = connector.Connect()
	if err != nil {
		t.Fatalf("Connect failure %v", err)
	}
	cs := connector.CoreSevice()
	p := compute.NewListInstancesParams().WithCompartmentID(cfg.Properties.OCI.CompartmentID).WithAvailabilityDomain(&ad)
	res, err := cs.Compute.ListInstances(p)
	if err != nil {
		t.Fatalf(" Error getting list of instances %v", err)
	}

	for _, in := range res.Payload {
		assertNotNil(t, in.ImageID, "Empty instance image ID")
	}
}
