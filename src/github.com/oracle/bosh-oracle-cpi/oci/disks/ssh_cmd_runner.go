package disks

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"errors"
	"fmt"
	"strings"
	"time"
)

type sshCmdRunner struct {
	logger       boshlog.Logger
	user         string
	remoteIP     string
	identityPath string
	runner       boshsys.CmdRunner
}

const cmdArgTemplate = "-o StrictHostKeyChecking=no %v@%v -i %s %s"

type SSHCmdResultHandler func(stdout string, stderr string) (retry bool, reasonToRetry string)

type SSHCmdRunner interface {
	RunCommand(cmd string, resultHandler SSHCmdResultHandler, maxAttempts int, durationBetweenAttempts time.Duration) error
	Connect(maxAttempts int, durationBetweenAttempts time.Duration) error
}

func NewSSHCmdRunner(user string, remoteIP string, identityPath string,
	logger boshlog.Logger) SSHCmdRunner {

	return sshCmdRunner{user: user, remoteIP: remoteIP, identityPath: identityPath,
		logger: logger, runner: boshsys.NewExecCmdRunner(logger),
	}
}

func (r sshCmdRunner) RunCommand(cmd string, resultHandler SSHCmdResultHandler,
	maxAttempts int, durationBetweenAttempts time.Duration) error {

	attemptFunc := func() (bool, error) {
		stdout, stderr, _, err := r.runner.RunComplexCommand(r.newCommand(cmd))
		if err != nil {
			r.logger.Debug(diskOperationsLogTag, "Error running cmd %s %v", cmd, err)
			return false, err
		}
		retry, reason := resultHandler(stdout, stderr)
		if retry {
			return false, errors.New(reason)
		} else {
			return true, nil
		}
	}
	return boshretry.NewAttemptRetryStrategy(maxAttempts, durationBetweenAttempts,
		boshretry.NewRetryable(attemptFunc),
		r.logger).Try()
}

func (r sshCmdRunner) Connect(maxAttempts int, durationBetweenAttempts time.Duration) error {

	attemptFunc := func() (bool, error) {
		_, _, _, err := r.runner.RunComplexCommand(r.newCommand("echo alive"))
		return err != nil, err
	}

	return boshretry.NewAttemptRetryStrategy(maxAttempts, durationBetweenAttempts,
		boshretry.NewRetryable(attemptFunc),
		r.logger).Try()

}

func (r sshCmdRunner) newCommand(cmd string) boshsys.Command {
	args := fmt.Sprintf(cmdArgTemplate, r.user, r.remoteIP, r.identityPath, cmd)
	return boshsys.Command{Name: "ssh", Args: strings.Split(args, " ")}
}
