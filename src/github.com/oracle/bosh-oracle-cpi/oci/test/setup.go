package test

import (
	"github.com/oracle/bosh-oracle-cpi/oci/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"os"
	"path/filepath"
	"strings"
)

type Fixture interface {
	Setup() error
	TearDown() error
}

const defaultTestSection string = "CPITEST"
const defaultIniPath = "~/.oci/config"

const volume1GB int64 = 1024

var vmStandard12config = vm.InstanceConfiguration{
	Shape: "VM.Standard1.2",
	Name:  "VM-test-fixture",
}

var manualNetworkNoIp = registry.NetworkSetting{
	Type:          "manual",
	IP:            "",
	Gateway:       "10.0.1.1",
	Netmask:       "255.255.255.0",
	UseDHCP:       true,
	DNS:           []string{"8.8.8.8"},
	Default:       []string{"dns", "gateway"},
	Resolved:      true,
	Preconfigured: true,
}

func iniConfigPath() string {
	path := os.Getenv("CPITEST_CONFIG")

	if path == "" {
		path = defaultIniPath
	}

	path = strings.Replace(path, "~", os.Getenv("HOME"), -1)

	if !filepath.IsAbs(path) {
		dir, _ := os.Getwd()
		path = filepath.Join(dir, path)

	}
	return path
}

func iniTestSection() string {
	section := os.Getenv("CPITEST_PROFILE")
	if section == "" {
		section = defaultTestSection
	}
	return section
}
