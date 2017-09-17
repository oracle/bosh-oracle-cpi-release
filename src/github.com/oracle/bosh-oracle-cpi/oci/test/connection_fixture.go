package test

import (
	"github.com/oracle/bosh-oracle-cpi/oci/client"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"testing"
)

type ConnectionFixture struct {
	connector  client.Connector
	logger     boshlog.Logger
	cpiTestIni CpiTestIni
}

func NewConnectionFixture() *ConnectionFixture {
	return &ConnectionFixture{}
}

func (c *ConnectionFixture) Setup(t *testing.T) error {
	cfg, ini, err := NewTestConfig(iniConfigPath(), iniTestSection())
	if err != nil {
		t.Fatalf("Error creating config %v", err)
		return err
	}
	logLevel := boshlog.LevelWarn
	if ini.logLevel != "" {
		logLevel, _ = boshlog.Levelify(ini.logLevel)
	}
	logger := boshlog.NewLogger(logLevel)
	connector := client.NewConnector(cfg, logger)
	err = connector.Connect()
	if err != nil {
		t.Fatalf("Connection failure  %v", err)
		return err
	}
	c.connector = connector
	c.logger = logger
	c.cpiTestIni = ini
	return nil
}

func (c *ConnectionFixture) TearDown(t *testing.T) error {
	return nil
}

func (c *ConnectionFixture) Connector() client.Connector {
	return c.connector
}

func (c *ConnectionFixture) Logger() boshlog.Logger {
	return c.logger
}

func (c *ConnectionFixture) VCN() string {
	return c.cpiTestIni.VcnName
}
func (c *ConnectionFixture) Subnet() string {
	return c.cpiTestIni.SubnetName
}

func (c *ConnectionFixture) AD() string {
	return c.cpiTestIni.AvailabilityDomain
}

func (c *ConnectionFixture) StemcellImageID() string {
	return c.cpiTestIni.StemcellImageID
}
