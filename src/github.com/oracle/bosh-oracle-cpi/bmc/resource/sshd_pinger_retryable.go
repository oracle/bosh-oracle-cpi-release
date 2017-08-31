package resource

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/retrystrategy"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"fmt"
)

type sshdChecker struct {
	cmdRunner    boshsys.CmdRunner
	cmd          boshsys.Command
	logger       boshlog.Logger
	remoteServer string
}

func NewSSHDCheckerRetryable(user string, remoteIP string, logger boshlog.Logger) retrystrategy.Retryable {
	userAtHost := fmt.Sprintf("%v@%v", user, remoteIP)

	args := []string{
		"-o", "StrictHostKeyChecking=no",
		userAtHost,
		"echo",
		"alive",
	}
	cmd := boshsys.Command{Name: "ssh", Args: args}
	return sshdChecker{cmdRunner: boshsys.NewExecCmdRunner(logger),
		cmd: cmd, logger: logger, remoteServer: remoteIP}

}

func (r sshdChecker) Attempt() (bool, error) {
	_, _, status, err := r.cmdRunner.RunComplexCommand(r.cmd)
	if err != nil {
		r.logger.Debug(logTag, "sshd possibly not up on %s", r.remoteServer)
		return true, err
	}
	r.logger.Debug(logTag, "ssh status=%v", status)
	return true, err
}
