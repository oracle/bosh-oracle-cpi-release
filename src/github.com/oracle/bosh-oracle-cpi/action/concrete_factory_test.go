package action_test

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	fakeuuid "github.com/cloudfoundry/bosh-utils/uuid/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/oracle/bosh-oracle-cpi/action"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	clientfakes "github.com/oracle/bosh-oracle-cpi/bmc/client/fakes"
	"github.com/oracle/bosh-oracle-cpi/config"
	"github.com/oracle/bosh-oracle-cpi/registry"
)

var _ = Describe("ConcreteFactory", func() {
	var (
		uuidGen        *fakeuuid.FakeGenerator
		connector      client.Connector
		logger         boshlog.Logger
		registryClient registry.Client

		cfg = config.Config{
			Cloud: config.Cloud{
				Properties: config.CPIProperties{
					Registry: registry.ClientOptions{
						Protocol: "http",
						Host:     "fake-host",
						Port:     5555,
						Username: "fake-username",
						Password: "fake-password",
					},
				},
			},
		}

		factory Factory
	)

	BeforeEach(func() {
		connector = &clientfakes.FakeConnector{}
		uuidGen = &fakeuuid.FakeGenerator{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		factory = NewConcreteFactory(
			connector,
			uuidGen,
			cfg,
			logger,
		)
	})

	BeforeEach(func() {
		registryClient = registry.NewHTTPClient(cfg.Cloud.Properties.Registry, logger)
	})

	It("returns error if action cannot be created", func() {
		action, err := factory.Create("fake-unknown-action")
		Expect(err).To(HaveOccurred())
		Expect(action).To(BeNil())
	})

	It("returns creates stemcell action", func() {
		action, err := factory.Create("create_stemcell")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewCreateStemcell(connector, logger)))
	})

	It("returns delete stemcell action", func() {
		action, err := factory.Create("delete_stemcell")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewDeleteStemcell(connector, logger)))
	})

	// VM actions
	It("returns create vm action", func() {
		action, err := factory.Create("create_vm")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewCreateVM(connector, logger, registryClient, uuidGen)))
	})

	It("returns delete vm action", func() {
		action, err := factory.Create("delete_vm")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewDeleteVM(connector, logger, registryClient)))
	})
	It("returns has vm action", func() {
		action, err := factory.Create("has_vm")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewHasVM(connector, logger)))
	})

	It("returns create disk action", func() {
		action, err := factory.Create("create_disk")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewCreateDisk(connector, logger)))
	})

	It("returns delete disk action", func() {
		action, err := factory.Create("delete_disk")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewDeleteDisk(connector, logger)))
	})

	It("returns attach disk action", func() {
		action, err := factory.Create("attach_disk")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewAttachDisk(connector, logger, registryClient)))
	})

	It("returns detach disk action", func() {
		action, err := factory.Create("detach_disk")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewDetachDisk(connector, logger, registryClient)))
	})
	It("returns has disk action", func() {
		action, err := factory.Create("has_disk")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewHasDisk(connector, logger)))
	})

	It("returns get disks action", func() {
		action, err := factory.Create("get_disks")
		Expect(err).ToNot(HaveOccurred())
		Expect(action).To(Equal(NewGetDisks(connector, logger)))
	})
})
