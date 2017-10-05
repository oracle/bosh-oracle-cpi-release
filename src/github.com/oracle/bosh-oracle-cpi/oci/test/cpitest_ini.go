package test

import (
	"github.com/go-ini/ini"
	"os"
	"strings"
)

// CpiTestIni holds the configuration
// used for running the CPI tests
// Default path is ~/.oci/config
type CpiTestIni struct {
	Tenant                 string `ini:"tenancy" mapstructure:"tenant"`
	User                   string `ini:"user" mapstructure:"user"`
	Fingerprint            string `ini:"fingerprint" mapstructure:"fingerprint"`
	Host                   string `ini:"host" mapstructure:"host"`
	KeyFile                string `ini:"key_file" mapstructure:"KeyFile"`
	CompartmentId          string `ini:"compartment" mapstructure:"CompartmentId"`
	Region                 string `ini:"region" mapstructure:"Region"`
	AvailabilityDomain     string `ini:"ad" mapstructure:"AvailabilityDomain"`
	VcnName                string `ini:"vcn" mapstructure:"VcnName"`
	SubnetName             string `ini:"subnet" mapstructure:"SubnetName"`
	CpiUser                string `ini:"cpiUser" mapstructure:"CpiUser"`
	CpiPrivateKeyPath      string `ini:"cpiPrivateKeyPath" mapstructure:"CpiPrivateKeyPath"`
	CpiPublicKeyPath       string `ini:"cpiPublicKeyPath" mapstructure:"CpiPublicKeyPath"`
	UserPublicKeyPath      string `ini:"userPublicKeyPath" mapstructure:"UserPublicKeyPath"`
	StemcellImageID        string `ini:"stemcellImage" mapstructure:"StemcellImageID"`
	StemcellImageSourceURI string `ini:"stemcellImageSourceURI" mapstructure:"StemcellImageSourceURI"`
	LogLevel               string `ini:"logLevel" mapstructure:"LogLevel"`
}

func NewCpiTestIni(filePath string, section string) (CpiTestIni, error) {

	cfg := CpiTestIni{}
	if filePath == "" {
		filePath = defaultIniPath
	}
	if section == "" {
		section = defaultTestSection
	}
	err := (&cfg).loadIni(filePath, section)

	return cfg, err

}

func (b *CpiTestIni) loadIni(filePath string, section string) error {

	path := strings.Replace(filePath, "~", os.Getenv("HOME"), -1)
	t, err := ini.Load(path)
	if err != nil {
		return err
	}
	s, err := t.GetSection(section)
	if err != nil {
		return err
	}
	err = s.MapTo(b)
	b.KeyFile = strings.Replace(b.KeyFile, "~", os.Getenv("HOME"), -1)
	return nil
}
