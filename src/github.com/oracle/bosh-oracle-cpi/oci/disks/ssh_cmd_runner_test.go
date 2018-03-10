package disks

import (
	"errors"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("SSHCmdRunner", func() {
	var (
		fakeRunner       = fakes.NewFakeCmdRunner()
		logger           = boshlog.NewLogger(boshlog.LevelNone)
		remoteUser       = "test-user"
		remoteIP         = "test-IP"
		userKeyPath      = "test-user-keypath"
		fullCmdTemplate  = fmt.Sprintf("ssh %s", cmdArgTemplate)
		sshCommandRunner = sshCmdRunner{logger, remoteUser, remoteIP,
			userKeyPath, fakeRunner}
	)
	Context("when remote command succeeds", func() {
		var (
			remoteCmd       = "echo Hello"
			remoteCmdOutput = "Hello"
			handlerCalled   bool
			outToHandler    string
			errToHandler    string
			err             error
		)
		BeforeEach(func() {
			fullCmd := fmt.Sprintf(fullCmdTemplate, remoteUser, remoteIP, userKeyPath, remoteCmd)
			cmdResultHandler := func(stdout, stderr string) (bool, string) {
				handlerCalled = true
				outToHandler = stdout
				errToHandler = stderr
				return false, ""
			}
			fakeRunner.AddCmdResult(fullCmd, fakes.FakeCmdResult{Stdout: remoteCmdOutput,
				Stderr: "", ExitStatus: 0, Error: nil, Sticky: false})
			err = sshCommandRunner.RunCommand(remoteCmd, cmdResultHandler, 1, time.Nanosecond)

		})
		It("doesn't return error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
		It("calls output handler", func() {
			Expect(handlerCalled).To(BeTrue())
		})
		It("passes remote stdout and stderr to handler", func() {
			Expect(outToHandler).To(Equal(remoteCmdOutput))
			Expect(errToHandler).To(Equal(""))
		})

	})
	Context("when remote command fails", func() {
		It("bubbles up the error when command fails", func() {

			remoteCommand := "non-existent-command"
			fullCmd := fmt.Sprintf(fullCmdTemplate, remoteUser, remoteIP, userKeyPath, remoteCommand)
			remoteCmdError := errors.New("Remote command Failed")
			fakeRunner.AddCmdResult(fullCmd, fakes.FakeCmdResult{Stdout: "",
				Stderr: "Unable to find command", ExitStatus: 1, Error: remoteCmdError, Sticky: false})

			var invoked bool
			cmdResultHandler := func(stdout, stderr string) (bool, string) {
				invoked = true
				return false, ""
			}
			err := sshCommandRunner.RunCommand(remoteCommand, cmdResultHandler, 2, time.Nanosecond)
			Expect(invoked).To(BeFalse())
			Expect(err).To(Equal(remoteCmdError))

		})

	})
})
