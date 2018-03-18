package disks

import (
	"errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("SSHCmdRunner", func() {
	var (
		logger           = boshlog.NewLogger(boshlog.LevelNone)
		remoteUser       = "test-user"
		remoteIP         = "test-IP"
		userKeyPath      = "test-user-keypath"
		sshCommandRunner = sshCmdRunner{logger, remoteUser, remoteIP,
			userKeyPath, nil}
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
			cmdResultHandler := func(stdout, stderr string) (bool, string) {
				handlerCalled = true
				outToHandler = stdout
				errToHandler = stderr
				return false, ""
			}
			err = sshCommandRunner.RunCommand(remoteCmd, cmdResultHandler, 1, time.Nanosecond)

		})
		XIt("doesn't return error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
		XIt("calls output handler", func() {
			Expect(handlerCalled).To(BeTrue())
		})
		XIt("passes remote stdout and stderr to handler", func() {
			Expect(outToHandler).To(Equal(remoteCmdOutput))
			Expect(errToHandler).To(Equal(""))
		})

	})
	Context("when remote command fails", func() {
		XIt("bubbles up the error when command fails", func() {

			remoteCommand := "non-existent-command"
			remoteCmdError := errors.New("Remote command Failed")
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
