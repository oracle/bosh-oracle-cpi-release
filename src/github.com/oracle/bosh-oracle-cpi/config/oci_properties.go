package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"oracle/oracle-iaas-go.git/transport"
)

const stemcellUserName string = "vcap"

// OCIProperties contains the properties for configuring
// BOSH CPI for Oracle Cloud Infrastructure
type OCIProperties struct {
	// Tenancy under which the resources are provisioned
	Tenancy string `json:"tenancy"`

	// User OCID
	User string `json:"user"`

	// Compartment OCID
	CompartmentID string `json:"compartment"`

	// Region name
	Region string `json:"region"`

	// Fingerprint of the User API key
	Fingerprint string `json:"fingerprint"`

	// APIKeyFile is the path to the private API key
	APIKeyFile string `json:"apikeyfile"`

	// CPIKeyfile is the path to the private key used by the CPI
	// used for SSH connections
	CpiKeyFile string `json:"cpikeyfile"`

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
func (b OCIProperties) Validate() error {

	if err := isAnyEmpty(map[string]string{
		"tenancy":     b.Tenancy,
		"user":        b.User,
		"fingerprint": b.Fingerprint,
		"apikeyfile":  b.APIKeyFile,
		"compartment": b.CompartmentID,
		"cpikeyfile":  b.CpiKeyFile,
	}); err != nil {
		return err
	}
	return validateFilePaths([]string{b.APIKeyFile})
}

func isAnyEmpty(attributes map[string]string) error {
	for name, value := range attributes {
		if value == "" {
			return fmt.Errorf(" Property %s must not be empty", name)
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

func newSanitizedConfig(configFullPath string, b OCIProperties) OCIProperties {
	dir := filepath.Dir(configFullPath)

	return OCIProperties{
		Tenancy:           b.Tenancy,
		User:              b.User,
		CompartmentID:     b.CompartmentID,
		Region:            b.Region,
		Fingerprint:       b.Fingerprint,
		APIKeyFile:        filepath.Join(dir, filepath.Base(b.APIKeyFile)),
		CpiKeyFile:        filepath.Join(dir, filepath.Base(b.CpiKeyFile)),
		UsePublicIPForSSH: b.UsePublicIPForSSH,
		AuthorizedKeys:    b.AuthorizedKeys,
		SSHTunnel:         b.SSHTunnel,
	}
}

// TransportConfig returns the configuration properties
// needed by the underlying transport layer for communicating
// with OCI
func (b OCIProperties) TransportConfig(host string) transport.Config {

	return transport.Config{Tenant: b.Tenancy, User: b.User,
		Fingerprint: b.Fingerprint, Host: host, KeyFile: b.APIKeyFile}
}

// UserSSHPublicKeyContent returns the configured ssh-rsa user public key
func (b OCIProperties) UserSSHPublicKeyContent() string {
	return sanitizeSSHKey(b.AuthorizedKeys.User)
}

// CpiSSHPublicKeyContent returns the configured cpi user's ssh-rsa public key
func (b OCIProperties) CpiSSHPublicKeyContent() string {
	return sanitizeSSHKey(b.AuthorizedKeys.Cpi)
}

// CpiSSHConfig returns the CPI ssh configuration
func (b OCIProperties) CpiSSHConfig() SSHConfig {
	return SSHConfig{stemcellUserName, b.CpiKeyFile, b.UsePublicIPForSSH}
}

func sanitizeSSHKey(key string) string {
	if key != "" {
		return strings.TrimSuffix(strings.TrimSpace(key), "\n")
	}
	return key
}

func (b OCIProperties) AuthorizedKeysContents() []string {
	keys := []string{}
	userKey := b.UserSSHPublicKeyContent()
	if userKey != "" {
		keys = append(keys, userKey)
	}
	cpiKey := b.CpiSSHPublicKeyContent()
	if cpiKey != "" {
		keys = append(keys, cpiKey)
	}
	return keys
}
