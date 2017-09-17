package resource

import (
	"time"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"

	"errors"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/baremetal/core/client/compute"
)

type Instance struct {
	ocid     string
	location Location

	// May not always be known
	publicIPs  []string
	privateIPs []string
}

func NewInstance(ocid string, loc Location) *Instance {
	return &Instance{ocid: ocid, location: loc}
}

func NewInstanceWithPrivateIPs(ocid string, loc Location, privateIPs []string) *Instance {
	return &Instance{ocid: ocid, location: loc, privateIPs: privateIPs}
}

func (in *Instance) ID() string {
	return in.ocid
}

func (in *Instance) EnsureReachable(c client.Connector, l boshlog.Logger) error {

	err := in.queryIPs(c, l)
	if err != nil {
		return err
	}
	return in.setupSSHTunnelToAgent(c, l)
}

func (in *Instance) queryIPs(c client.Connector, l boshlog.Logger) error {
	err := in.waitUntilStarted(c, l)
	if err != nil {
		return err
	}

	var public []string
	var private []string
	public, private, err = in.location.instanceIPs(c, in.ocid)
	if err != nil {
		l.Debug(logTag, "Error finding IPs %s", err)
		return err
	}

	in.publicIPs = make([]string, len(public))
	in.privateIPs = make([]string, len(private))
	copy(in.publicIPs, public)
	copy(in.privateIPs, private)

	l.Debug(logTag, "Queried IPs, Private %v, Public %v", in.privateIPs, in.publicIPs)
	return nil
}

func (in *Instance) waitUntilStarted(c client.Connector, l boshlog.Logger) (err error) {

	getInstanceState := func() (bool, error) {

		p := compute.NewGetInstanceParams().WithInstanceID(in.ocid)
		r, err := c.CoreSevice().Compute.GetInstance(p)
		if err != nil {
			return false, err
		}
		instance := r.Payload
		switch *instance.LifecycleState {
		case "PROVISIONING", "CREATING_IMAGE":
			return true, errors.New("Not provisioned yet")
		case "STOPPING", "STOPPED", "TERMINATING", "TERMINATED":
			return false, errors.New("Terminated")
		case "RUNNING":
			return true, nil
		case "STARTING":
			return true, errors.New("Starting")
		}
		return true, errors.New("Unknown instance state")
	}

	retryable := boshretry.NewRetryable(getInstanceState)
	retryStrategy := boshretry.NewUnlimitedRetryStrategy(10*time.Second, retryable, l)

	l.Debug(logTag, "Waiting for instance to reach RUNNING state...")
	if err := retryStrategy.Try(); err != nil {
		l.Debug(logTag, "Error waiting for instance to start %v", err)
		return err
	}
	l.Debug(logTag, "Done")
	return nil
}

func (in *Instance) PublicIP(c client.Connector, l boshlog.Logger) (string, error) {
	if !in.havePublicIP() {
		if err := in.queryIPs(c, l); err != nil {
			return "", err
		}
	}
	return in.publicIPs[0], nil
}

func (in *Instance) PrivateIP(c client.Connector, l boshlog.Logger) (string, error) {
	if !in.havePrivateIP() {
		if err := in.queryIPs(c, l); err != nil {
			return "", err
		}
	}
	return in.privateIPs[0], nil
}

func (in *Instance) setupSSHTunnelToAgent(c client.Connector, l boshlog.Logger) (err error) {
	tunnel := c.SSHTunnelConfig()
	if tunnel.IsConfigured() {

		duration, _ := time.ParseDuration(tunnel.Duration)
		remotePort, _ := c.AgentOptions().MBusPort()
		remoteIP, err := in.remoteIP(c, l, tunnel.UsePublicIP)
		if err != nil {
			return err
		}

		// Ensure SSHD is up
		retryable := NewSSHDCheckerRetryable(tunnel.User, remoteIP, l)
		strategy := boshretry.NewAttemptRetryStrategy(10, 20*time.Second, retryable, l)
		err = strategy.Try()

		// Then start the port forwarder
		if err == nil {
			retryable = NewSSHPortForwarderRetryable(tunnel.LocalPort, remotePort, remoteIP, tunnel.User,
				duration, l)
			strategy = boshretry.NewAttemptRetryStrategy(2, 2*time.Second, retryable, l)
			err = strategy.Try()
		}
		return err
	}
	return nil
}

func (in *Instance) havePublicIP() bool {
	return in.publicIPs != nil && len(in.publicIPs) > 0
}

func (in *Instance) havePrivateIP() bool {
	return in.privateIPs != nil && len(in.privateIPs) > 0
}

func (in *Instance) remoteIP(c client.Connector, l boshlog.Logger, public bool) (string, error) {
	if public {
		return in.PublicIP(c, l)
	} else {
		return in.PrivateIP(c, l)
	}

}

func (in *Instance) Location() Location {
	return in.location
}
