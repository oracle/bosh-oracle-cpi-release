package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"oracle/oracle-iaas-go.git/transport"
)

// BmcsProperties contains the properties for configuring
// Oracle BMC CPI
type BmcsProperties struct {
	// Tenancy under which the resources are provisioned
	Tenancy string `json:"tenancy"`

	// User OCID
	User string `json:"user"`

	// Compartment OCID
	CompartmentID string `json:"compartment"`

	// Region name
	Region string `json:"region, omitempty"`

	// Fingerprint of the User API key
	Fingerprint string `json:"fingerprint"`

	// APIKeyFile is the path to the private API key
	APIKeyFile string `json:"apikeyfile"`

	// CPIKeyfile is the path to the private key used by the CPI
	// used for SSH connections
	CpiKeyFile string `json:"cpikeyfile"`

	// CpiUser is name of the user to use for CPI SSH connections
	CpiUser string `json:"cpiuser"`

	// UsePublicIPForSSH controls whether to use public or private IP
	// of the target insatnce for establishing SSH connections
	UsePublicIPForSSH bool `json:"usePublicIpForSsh,omitempty"`

	// AuthorizedKeys contains the public ssh-keys to provision
	// on new vms
	AuthorizedKeys AuthorizedKeys `json:"authorized_keys"`

	// SSHTunnel is the configuration for creating a forward SSH tunnel
	SSHTunnel SSHTunnel `json:"sshTunnel,omitempty"`
}

// AuthorizedKeys is the set of public
// ssh-rsa keys to be installed
// on the default initial account
// provisioned on a new vm
type AuthorizedKeys struct {
	Cpi  string `json:"cpi"`
	User string `json:"user, omitempty"`
}

// Validate raises an error if any of the mandatory
// properties are missing
func (b BmcsProperties) Validate() error {

	if err := isAnyEmpty([]string{
		b.Tenancy,
		b.User,
		b.Fingerprint,
		b.APIKeyFile,
		b.CompartmentID,
		b.CpiUser,
	}); err != nil {
		return err
	}
	return validateFilePaths([]string{b.APIKeyFile})
}

func isAnyEmpty(attributes []string) error {
	for _, s := range attributes {
		if s == "" {
			return fmt.Errorf(" bmcs.%s must be non-empty", s)
		}
	}
	return nil
}

func validateFilePaths(paths []string) error {
	for _, path := range paths {
		if err := validateFilePath(path); err != nil {
			return err
		}
	}
	return nil
}

func validateFilePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("File %s doesn't exist", path)
	}
	return nil
}

func newSanitizedConfig(configFullPath string, b BmcsProperties) BmcsProperties {
	dir := filepath.Dir(configFullPath)

	return BmcsProperties{
		Tenancy:           b.Tenancy,
		User:              b.User,
		CompartmentID:     b.CompartmentID,
		Region:            b.Region,
		Fingerprint:       b.Fingerprint,
		APIKeyFile:        filepath.Join(dir, filepath.Base(b.APIKeyFile)),
		CpiKeyFile:        filepath.Join(dir, filepath.Base(b.CpiKeyFile)),
		CpiUser:           b.CpiUser,
		UsePublicIPForSSH: b.UsePublicIPForSSH,
		AuthorizedKeys:    b.AuthorizedKeys,
		SSHTunnel:         b.SSHTunnel,
	}
}

// TransportConfig returns the configuration properties
// needed by the underlying transport layer for communicating
// with BMC
func (b BmcsProperties) TransportConfig(host string) transport.Config {

	return transport.Config{Tenant: b.Tenancy, User: b.User,
		Fingerprint: b.Fingerprint, Host: host, KeyFile: b.APIKeyFile}
}

// UserSSHPublicKeyContent returns the configured ssh-rsa user public key
func (b BmcsProperties) UserSSHPublicKeyContent() (string, error) {
	return sanitizeSSHKey(b.AuthorizedKeys.User)
}

// CpiSSHPublicKeyContent returns the configured cpi user's ssh-rsa public key
func (b BmcsProperties) CpiSSHPublicKeyContent() (string, error) {
	return sanitizeSSHKey(b.AuthorizedKeys.Cpi)
}

// CpiSSHConfig returns the CPI ssh configuration
func (b BmcsProperties) CpiSSHConfig() SSHConfig {
	return SSHConfig{b.CpiUser, b.CpiKeyFile, b.UsePublicIPForSSH}
}

func sanitizeSSHKey(key string) (string, error) {
	return strings.TrimSuffix(strings.TrimSpace(key), "\n"), nil
}
