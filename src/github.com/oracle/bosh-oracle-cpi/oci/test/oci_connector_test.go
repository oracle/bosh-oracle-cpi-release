package test

import (
	"github.com/oracle/bosh-oracle-cpi/config"
	"github.com/oracle/bosh-oracle-cpi/oci/client"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"oracle/oci/core/client/compute"
	"path/filepath"
	"strings"
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

func Test_SingleSSHKeyFormattedCorrectly(t *testing.T) {

	logger := boshlog.NewLogger(boshlog.LevelNone)
	fs := boshsys.NewOsFileSystem(logger)
	json := filepath.Join(assetsDir(), "cpi_without_userauthorizedkey.json")

	c, err := config.NewConfigFromPath(json, fs)
	if err != nil {
		t.Fatalf(" Error building config %v", err)
	}
	connector := client.NewConnector(c.Cloud, logger)

	keys := connector.AuthorizedKeys()
	if len(keys) != 1 {
		t.Logf("Unexpected number of keys. Expected 1 Actual %v", len(keys))
		t.Fail()
	}
	failIfHasNewlineChars(t, keys[0])
}

func Test_TwoSSHKeyFormattedCorrectly(t *testing.T) {

	logger := boshlog.NewLogger(boshlog.LevelNone)
	fs := boshsys.NewOsFileSystem(logger)
	json := filepath.Join(assetsDir(), "cpi_with_userauthorizedkey.json")

	c, err := config.NewConfigFromPath(json, fs)
	if err != nil {
		t.Fatalf(" Error building config: %v", err)
		t.Fail()
	}
	connector := client.NewConnector(c.Cloud, logger)

	keys := connector.AuthorizedKeys()
	if len(keys) != 2 {
		t.Logf("Unexpected number of keys. Expected 1 Actual %v", len(keys))
		t.Fail()
	}
	for _, k := range keys {
		failIfHasNewlineChars(t, k)
	}
}

func failIfHasNewlineChars(t *testing.T, key string) {

	if strings.HasPrefix(key, "\n") {
		t.Logf("Key %s unexpectedly starts with newline", key)
		t.Fail()
	}
	if strings.HasSuffix(key, "\n") {
		t.Logf("Key %s unexpectedly ends with newline", key)
		t.Fail()
	}
}
