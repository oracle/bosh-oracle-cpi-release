package action

import (
	"errors"
	"fmt"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"github.com/oracle/bosh-oracle-cpi/oci/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"

	fakeuuid "github.com/cloudfoundry/bosh-utils/uuid/fakes"
	clientfakes "github.com/oracle/bosh-oracle-cpi/oci/client/fakes"
	vmfakes "github.com/oracle/bosh-oracle-cpi/oci/vm/fakes"
	registryfakes "github.com/oracle/bosh-oracle-cpi/registry/fakes"
)

var _ = Describe("CreateVM", func() {
	var (
		connector      *clientfakes.FakeConnector
		logger         boshlog.Logger
		registryClient *registryfakes.FakeClient
		uuidGen        *fakeuuid.FakeGenerator
		creator        *vmfakes.FakeVMCreator

		err             error
		networks        Networks
		cloudProps      VMCloudProperties
		env             Environment
		registryOptions registry.ClientOptions
		agentOptions    registry.AgentOptions

		expectedAgentSettings registry.AgentSettings
		vmCID                 VMCID

		createVM CreateVM
	)

	BeforeEach(func() {
		installVMCreatorFactory(func(c client.Connector, l boshlog.Logger, availabilityDomain string) vm.Creator {
			return creator
		})

		connector = &clientfakes.FakeConnector{}
		uuidGen = &fakeuuid.FakeGenerator{}
		uuidGen.GeneratedUUID = "fake-uuid"
		logger = boshlog.NewLogger(boshlog.LevelNone)
		registryClient = &registryfakes.FakeClient{}
		registryOptions = registry.ClientOptions{
			Protocol: "http",
			Host:     "fake-registry-host",
			Port:     25777,
			Username: "fake-registry-username",
			Password: "fake-registry-password",
		}
		agentOptions = registry.AgentOptions{
			Mbus: "http://fake-mbus",
			Blobstore: registry.BlobstoreOptions{
				Provider: "fake-blobstore-type",
			},
		}
		connector.AgentOptionsResult = agentOptions
		createVM = NewCreateVM(connector, logger, registryClient, uuidGen)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		BeforeEach(func() {

			cloudProps = VMCloudProperties{
				Shape:              "fake-BM-32",
				AvailabilityDomain: "fake-availabilityDomain",
			}
			networks = Networks{
				"fake-network-name": &Network{
					Type:    "manual",
					IP:      "10.0.0.X",
					Gateway: "fake-network-gateway",
					Netmask: "fake-network-netmask",
					DNS:     []string{"fake-network-dns"},
					DHCP:    true,
					Default: []string{"fake-network-default"},
					CloudProperties: NetworkCloudProperties{
						VcnName:    "fake-vcn",
						SubnetName: "fake-subnet1",
					},
				},
			}
		})
		Context("when no errors in vm creation", func() {
			BeforeEach(func() {
				creator = &vmfakes.FakeVMCreator{
					CreateInstanceResult: resource.NewInstance("fake-ocid", resource.Location{}),
					CreateInstanceError:  nil,
				}
			})
			It("creates the vm", func() {
				vmCID, err = createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps, networks, nil, env)
				Expect(err).NotTo(HaveOccurred())
				Expect(creator.CreateInstanceCalled).To(BeTrue())
				Expect(vmCID).To(Equal(VMCID("fake-ocid")))
			})
			It("uses uuid as part of vm name", func() {
				vmCID, err = createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps, networks, nil, env)
				Expect(creator.Configuration.Name).To(ContainSubstring(uuidGen.GeneratedUUID))
			})
			It("updates the registry", func() {
				expectedAgentSettings = registry.AgentSettings{
					AgentID: "fake-agent-id",
					Blobstore: registry.BlobstoreSettings{
						Provider: "fake-blobstore-type",
					},
					Disks: registry.DisksSettings{
						System:     "/dev/sda",
						Ephemeral:  "/dev/sda",
						Persistent: map[string]registry.PersistentSettings{},
					},
					Mbus: "http://fake-mbus",
					Networks: registry.NetworksSettings{
						"fake-network-name": registry.NetworkSetting{
							Type:          "manual",
							IP:            "10.0.0.X",
							Gateway:       "fake-network-gateway",
							Netmask:       "fake-network-netmask",
							DNS:           []string{"fake-network-dns"},
							UseDHCP:       true,
							Default:       []string{"fake-network-default"},
							Preconfigured: true,
						},
					},
					VM: registry.VMSettings{
						Name: fmt.Sprintf("bosh-%s", uuidGen.GeneratedUUID),
					},
				}
				vmCID, err = createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps, networks, nil, env)
				Expect(err).NotTo(HaveOccurred())
				Expect(registryClient.UpdateCalled).To(BeTrue())
				Expect(registryClient.UpdateSettings).To(Equal(expectedAgentSettings))
				Expect(vmCID).To(Equal(VMCID("fake-ocid")))
			})
		})
		Context("when vm creation fails", func() {
			BeforeEach(func() {
				creator = &vmfakes.FakeVMCreator{
					CreateInstanceResult: nil,
					CreateInstanceError:  errors.New("fake-launchinstance-error"),
				}
			})
			It("propagates the error", func() {
				vmCID, err = createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps, networks, nil, env)
				Expect(err).To(HaveOccurred())
				Expect(creator.CreateInstanceCalled).To(BeTrue())
				Expect(err.Error()).To(ContainSubstring("fake-launchinstance-error"))
				Expect(vmCID).To(Equal(VMCID("")))
			})
			It("doesn't update registry", func() {
				vmCID, err = createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps, networks, nil, env)
				Expect(err).To(HaveOccurred())
				Expect(creator.CreateInstanceCalled).To(BeTrue())
				Expect(registryClient.UpdateCalled).To(BeFalse())
				Expect(vmCID).To(Equal(VMCID("")))
			})
		})

	})

})
