package disks

import (
	"fmt"
	"time"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeCmdResult struct {
	stdout string
	stderr string
}

type fakeSSHCommandRunner struct {
	commandsRun   []string
	cmdResult     map[string]fakeCmdResult
	connectCalled bool
}

func (r *fakeSSHCommandRunner) RunCommand(cmd string, handler SSHCmdResultHandler, _ int, _ time.Duration) error {
	r.commandsRun = append(r.commandsRun, cmd)

	res, found := r.cmdResult[cmd]
	if found {
		handler(res.stdout, res.stderr)
	}
	return nil
}
func (r *fakeSSHCommandRunner) Connect(_ int, _ time.Duration) error {
	r.connectCalled = true
	return nil
}

func (r *fakeSSHCommandRunner) AddCommandResult(cmd string, stdout string, stderr string) {
	r.cmdResult[cmd] = fakeCmdResult{stdout: stdout, stderr: stderr}
}

var _ = Describe("RemoteISCSINodeAdministrator", func() {
	var (
		administrator remoteIscsiNodeAdministrator
		fakeRunner    *fakeSSHCommandRunner
		iqn                 = "fake-iqn"
		server              = "fake-serverip"
		port          int64 = 1024
		expected            = []string{
			fmt.Sprintf("sudo iscsiadm -m node -o new -T %s -p %s:%d", iqn, server, port),
			fmt.Sprintf("sudo iscsiadm -m node -T %s -p %s:%d -l", iqn, server, port),
			fmt.Sprintf("sudo iscsiadm -m node -o update -T %s -n node.startup -v automatic", iqn),
			fmt.Sprintf("lsscsi -t | grep disk |  grep %s | awk '{print $4}'", iqn),
		}
	)
	Context("when RunAttachment is called", func() {
		var (
			resolvedPath string
			err          error
		)
		BeforeEach(func() {
			fakeRunner = &fakeSSHCommandRunner{[]string{},
				map[string]fakeCmdResult{}, false}
			fakeRunner.AddCommandResult(expected[3], "/dev/test-block-device", "")
			administrator = remoteIscsiNodeAdministrator{
				logger:    boshlog.NewLogger(boshlog.LevelDebug),
				cmdRunner: fakeRunner}
			resolvedPath, err = administrator.RunAttachmentCommands(iqn, server, port)
		})
		It("connects to ssh server first", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(fakeRunner.connectCalled).To(Equal(true))
		})
		It("invokes multiple ssh commands", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(fakeRunner.commandsRun).To(Equal(expected))
		})
		It("returns resolved device path", func() {
			Expect(resolvedPath).To(Equal("/dev/test-block-device"))
			Expect(administrator.LastResolvedPath()).To(Equal("/dev/test-block-device"))
		})

	})
	Context("when lsscsi returns spaces or newline in device path", func() {
		var (
			resolvedPath string
			err          error
		)
		BeforeEach(func() {
			invalidDevicePath := fmt.Sprintf("/dev/fake-device     \n")
			fakeRunner = &fakeSSHCommandRunner{[]string{},
				map[string]fakeCmdResult{}, false}

			fakeRunner.AddCommandResult(expected[3], invalidDevicePath, "")
			administrator = remoteIscsiNodeAdministrator{
				logger:    boshlog.NewLogger(boshlog.LevelDebug),
				cmdRunner: fakeRunner}
			resolvedPath, err = administrator.RunAttachmentCommands(iqn, server, port)
		})
		It("trims the devicepath correctly", func() {
			Expect(resolvedPath).To(Equal("/dev/fake-device"))
			Expect(administrator.LastResolvedPath()).To(Equal("/dev/fake-device"))
		})

	})

})
