package disks

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"

	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"time"
)

type sshCmdRunner struct {
	logger       boshlog.Logger
	user         string
	remoteIP     string
	identityPath string

	client *ssh.Client
}

const cmdArgTemplate = "-vvv -o StrictHostKeyChecking=no %v@%v -i %s %s"
const remoteSSHPort = 22

type SSHCmdResultHandler func(stdout string, stderr string) (retry bool, reasonToRetry string)

type SSHCmdRunner interface {
	RunCommand(cmd string, resultHandler SSHCmdResultHandler, maxAttempts int, durationBetweenAttempts time.Duration) error
	Connect(maxAttempts int, durationBetweenAttempts time.Duration) error
}

func NewSSHCmdRunner(user string, remoteIP string, identityPath string,
	logger boshlog.Logger) SSHCmdRunner {

	return &sshCmdRunner{user: user, remoteIP: remoteIP, identityPath: identityPath, logger: logger}
}

func (r *sshCmdRunner) RunCommand(cmd string, resultHandler SSHCmdResultHandler,
	maxAttempts int, durationBetweenAttempts time.Duration) error {

	if r.client == nil {
		return fmt.Errorf("Not connected to %s", r.remoteIP)
	}

	attemptFunc := func() (bool, error) {
		r.logger.Debug(diskOperationsLogTag, "Attempting remote command %s", cmd)
		stdout, stderr, err := r.runRemoteCommand(cmd)
		if err != nil {
			r.logger.Debug(diskOperationsLogTag, "Failed with %v. Retrying", err)
			return true, err
		}
		retry, reason := resultHandler(stdout, stderr)
		if retry {
			return true, errors.New(reason)
		} else {
			r.logger.Debug(diskOperationsLogTag, "Success")
			return false, nil
		}
	}
	return boshretry.NewAttemptRetryStrategy(maxAttempts, durationBetweenAttempts,
		boshretry.NewRetryable(attemptFunc),
		r.logger).Try()
}

func (r *sshCmdRunner) Connect(maxAttempts int, durationBetweenAttempts time.Duration) error {

	// Read private key
	key, err := ioutil.ReadFile(r.identityPath)
	if err != nil {
		r.logger.Error(diskOperationsLogTag, "unable to read private key: %v", err)
		return err
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		r.logger.Error(diskOperationsLogTag, "unable to parse private key: %v", err)
		return err
	}

	config := &ssh.ClientConfig{
		User: r.user,
		Auth: []ssh.AuthMethod{
			ssh.RetryableAuthMethod(ssh.PublicKeys(signer), maxAttempts),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshEndPoint := fmt.Sprintf("%s:%d", r.remoteIP, remoteSSHPort)

	attemptFunc := func() (bool, error) {
		r.logger.Debug(diskOperationsLogTag, "Dialing %s", sshEndPoint)
		r.client, err = ssh.Dial("tcp", sshEndPoint, config)
		return err != nil, err
	}

	return boshretry.NewAttemptRetryStrategy(maxAttempts, durationBetweenAttempts,
		boshretry.NewRetryable(attemptFunc),
		r.logger).Try()
}

func (r *sshCmdRunner) runRemoteCommand(cmd string) (stdout string, stderr string, err error) {

	stdoutWriter := bytes.NewBufferString("")
	stderrWriter := bytes.NewBufferString("")

	session, err := r.client.NewSession()
	defer func() { session.Close() }()
	if err != nil {
		return "", "", err
	}

	// Issue command
	session.Stderr = stderrWriter
	session.Stdout = stdoutWriter
	if err := session.Start(cmd); err != nil {
		return "", "", err
	}

	// Wait for it to finish
	err = session.Wait()
	if err != nil {
		return "", "", err
	}

	// Resulting stdout/stderr
	return string(stdoutWriter.Bytes()), string(stderrWriter.Bytes()), nil
}
