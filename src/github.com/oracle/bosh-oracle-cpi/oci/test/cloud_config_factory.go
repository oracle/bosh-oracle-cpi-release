package test

import (
	"github.com/oracle/bosh-oracle-cpi/config"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func NewTestConfig(iniPath string, section string) (config.Cloud, CpiTestIni, error) {

	ini, err := NewCpiTestIni(iniPath, section)

	if err != nil {
		return config.Cloud{}, CpiTestIni{}, err
	}

	ociProperties := config.OCIProperties{
		Tenancy:           ini.Tenant,
		User:              ini.User,
		CompartmentID:     ini.CompartmentId,
		Region:            ini.Region,
		Fingerprint:       ini.Fingerprint,
		APIKeyFile:        absolutePath(iniPath, ini.KeyFile),
		CpiKeyFile:        absolutePath(iniPath, ini.CpiPrivateKeyPath),
		UsePublicIPForSSH: ini.UsePublicIPForSSH,
		AuthorizedKeys: config.AuthorizedKeys{
			User: publicKeyFileContent(absolutePath(iniPath, ini.UserPublicKeyPath)),
			Cpi:  publicKeyFileContent(absolutePath(iniPath, ini.CpiPublicKeyPath)),
		},
	}
	if err := ociProperties.Validate(); err != nil {
		return config.Cloud{}, CpiTestIni{}, err
	}
	cpiProps := config.CPIProperties{
		OCI:      ociProperties,
		Agent:    registry.AgentOptions{},
		Registry: registry.ClientOptions{},
	}
	return config.Cloud{Plugin: "oracle", Properties: cpiProps}, ini, nil

}

func publicKeyFileContent(path string) string {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(key)
}

func absolutePath(iniPath string, relativePath string) string {

	if strings.Contains(relativePath, "~") {
		return strings.Replace(relativePath, "~", os.Getenv("HOME"), -1)
	}
	basePath := filepath.Dir(iniPath)
	return filepath.Join(basePath, relativePath)
}

func assetsDir() string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, "assets")
}
