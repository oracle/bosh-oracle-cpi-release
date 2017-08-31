package config

// SSHConfig contains configuration used
// by the CPI for establishing SSH connection
// for configuring new vm instances
type SSHConfig struct {
	remoteUser           string
	localIdentityKeyPath string
	usePublicIP          bool
}

// RemoteUser returns the name of the user
// on the target instance
func (c SSHConfig) RemoteUser() string {
	return c.remoteUser
}

// LocalIdentityKeyPath returns the path to
// the identity file containing the key
// for establishing the connection
func (c SSHConfig) LocalIdentityKeyPath() string {
	return c.localIdentityKeyPath
}

// UsePublicIP returns true if the SSH connection
// should be tried on the public IP of the
// target instance
func (c SSHConfig) UsePublicIP() bool {
	return c.usePublicIP
}
