package disks

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"strings"
	"time"
)

type IscsiNodeAdministrator interface {
	RunAttachmentCommands(scsiIQN string, scsiServer string,
		scsiServerPort int64) (resolvedDevicePath string, err error)
}

const addCmd = "sudo iscsiadm -m node -o new -T %s -p %s:%d"
const loginCmd = "sudo iscsiadm -m node -T %s -p %s:%d -l"
const persistCmd = "sudo iscsiadm -m node -o update -T %s -n node.startup -v automatic"
const devicePath = "lsscsi -t | grep disk |  grep %s | awk '{print $4}'"

type scsiCommand struct {
	cmd                  string
	maxAttempts          int
	delayBetweenAttempts time.Duration
	resultHandler        SSHCmdResultHandler
}

type remoteIscsiNodeAdministrator struct {
	logger    boshlog.Logger
	cmdRunner SSHCmdRunner

	resolvedDevicePath string
}
type IscsciNodeAdministratorFactory func(remoteUser string, remoteIP string, localIdentityFilePath string,
	logger boshlog.Logger) IscsiNodeAdministrator

func NewRemoteIscsiNodeAdministrator(remoteUser string, remoteIP string,
	localIdentityFilePath string, logger boshlog.Logger) IscsiNodeAdministrator {

	return &remoteIscsiNodeAdministrator{
		cmdRunner: NewSSHCmdRunner(remoteUser, remoteIP, localIdentityFilePath, logger),
		logger:    logger,
	}
}

func (adm *remoteIscsiNodeAdministrator) RunAttachmentCommands(scsiIQN string, scsiServer string,
	scsiServerPort int64) (string, error) {

	if err := adm.cmdRunner.Connect(60, 1*time.Second); err != nil {
		return "", err
	}
	adm.resetResolvedPath()
	cmds := []scsiCommand{
		{fmt.Sprintf(addCmd, scsiIQN, scsiServer, scsiServerPort), 4, 0 * time.Second, adm.defaultHandler},
		{fmt.Sprintf(loginCmd, scsiIQN, scsiServer, scsiServerPort), 5, 2 * time.Second, adm.defaultHandler},
		{fmt.Sprintf(persistCmd, scsiIQN), 1, 0 * time.Second, adm.defaultHandler},
		{fmt.Sprintf(devicePath, scsiIQN), 1, 0 * time.Second, adm.devicePathReader},
	}
	for _, c := range cmds {
		if err := adm.runCommand(c); err != nil {
			adm.logger.Error(diskOperationsLogTag, "Failed to run command %s", c.cmd)
			return "", err
		}
	}
	return adm.LastResolvedPath(), nil
}

func (adm *remoteIscsiNodeAdministrator) runCommand(c scsiCommand) error {
	return adm.cmdRunner.RunCommand(c.cmd, c.resultHandler, c.maxAttempts, c.delayBetweenAttempts)
}

func (adm *remoteIscsiNodeAdministrator) defaultHandler(stdout string,
	stderr string) (retry bool, reasonToRetry string) {
	return false, ""
}

func (adm *remoteIscsiNodeAdministrator) devicePathReader(stdout string,
	stderr string) (retry bool, reasonToRetry string) {

	if stdout == "" {
		return true, "Device path unresolved"
	}
	if !strings.HasPrefix(stdout, "/dev/") {
		return false, fmt.Sprintf("Unexpected output %s", stdout)
	}
	adm.setResolvedPath(strings.TrimSpace(strings.TrimSuffix(stdout, "\n")))
	return false, ""
}

func (adm *remoteIscsiNodeAdministrator) resetResolvedPath() {
	adm.resolvedDevicePath = ""
}

func (adm *remoteIscsiNodeAdministrator) LastResolvedPath() string {
	return adm.resolvedDevicePath
}

func (adm *remoteIscsiNodeAdministrator) setResolvedPath(path string) {
	adm.resolvedDevicePath = path
}
