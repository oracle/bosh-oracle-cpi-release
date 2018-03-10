package resource

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/retrystrategy"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"fmt"
	"time"
)

type portForwarderRetryable struct {
	cmdRunner boshsys.CmdRunner
	cmd       boshsys.Command
	logger    boshlog.Logger
}

func NewSSHPortForwarderRetryable(localPort int, remotePort int, remoteIP string,
	user string, duration time.Duration, logger boshlog.Logger) retrystrategy.Retryable {

	portForward := fmt.Sprintf("127.0.0.1:%v:%v:%v", localPort, remoteIP, remotePort)
	userAtHost := fmt.Sprintf("%v@%v", user, remoteIP)
	durationInSecs := fmt.Sprintf("%v", duration.Seconds())

	args := []string{
		"-f",
		"-L", portForward,
		"-o", "StrictHostKeyChecking=no",
		userAtHost,
		"sleep",
		durationInSecs,
	}

	cmd := boshsys.Command{Name: "ssh", Args: args, KeepAttached: true}
	return portForwarderRetryable{cmdRunner: boshsys.NewExecCmdRunner(logger),
		cmd: cmd, logger: logger}

}

func (r portForwarderRetryable) Attempt() (bool, error) {
	_, err := r.cmdRunner.RunComplexCommandAsync(r.cmd)
	return true, err
}
