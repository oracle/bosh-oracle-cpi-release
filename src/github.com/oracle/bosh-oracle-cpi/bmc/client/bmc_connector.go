package client

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/config"
	"github.com/oracle/bosh-oracle-cpi/registry"
	cclient "oracle/baremetal/core/client"
	iclient "oracle/baremetal/identity/client"

	"oracle/oracle-iaas-go.git/transport"

	rclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	logTag                 = "BmcConnector"
	apiBasePath            = "20160918"
	coreServices           = "iaas.us-phoenix-1.oraclecloud.com"
	identityService string = "identity.us-phoenix-1.oraclecloud.com"
	scheme          string = "https"
)

type Connector interface {
	Connect() error
	CoreSevice() *cclient.CoreServices
	Tenancy() string
	CompartmentId() string
	AuthorizedKeys() []string
	AgentOptions() registry.AgentOptions
	AgentRegistryEndpoint() string
	SSHTunnelConfig() config.SSHTunnel
	SSHConfig() config.SSHConfig
}

type connectorImpl struct {
	config config.Cloud
	logger boshlog.Logger

	coreService *cclient.CoreServices
	iamService  *iclient.IdentityAndAccessManagementService
}

func NewConnector(c config.Cloud, logger boshlog.Logger) Connector {

	return &connectorImpl{config: c, logger: logger,
		coreService: nil, iamService: nil}
}

func (c *connectorImpl) Connect() error {
	return c.createServiceClients(c.config.Properties.Bmcs, apiBasePath)
}

func (c *connectorImpl) CoreSevice() *cclient.CoreServices {

	return c.coreService
}

func (c *connectorImpl) IamService() *iclient.IdentityAndAccessManagementService {
	return c.iamService
}

func (c *connectorImpl) Tenancy() string {
	return c.config.Properties.Bmcs.Tenancy
}

func (c *connectorImpl) CompartmentId() string {
	return c.config.Properties.Bmcs.CompartmentID
}

func (c *connectorImpl) AuthorizedKeys() []string {
	keys := []string{}
	userKey, err := c.config.Properties.Bmcs.UserSSHPublicKeyContent()
	if err != nil {
		c.logger.Debug(logTag, "Ignored error while getting user key %v", err)
	} else {
		keys = append(keys, userKey)
	}

	cpiKey, err := c.config.Properties.Bmcs.CpiSSHPublicKeyContent()
	if err != nil {
		c.logger.Debug(logTag, "Ignored error while getting cpi key %v", err)
	} else {
		keys = append(keys, cpiKey)

	}
	c.logger.Debug(logTag, "Authorized keys %v", keys)
	return keys
}

func (c *connectorImpl) AgentOptions() registry.AgentOptions {
	return c.config.Properties.Agent
}

func (c *connectorImpl) AgentRegistryEndpoint() string {
	return c.config.Properties.Registry.EndpointWithCredentials()
}

func (c *connectorImpl) SSHTunnelConfig() config.SSHTunnel {
	return c.config.Properties.Bmcs.SSHTunnel
}

func (c *connectorImpl) SSHConfig() config.SSHConfig {
	return c.config.Properties.Bmcs.CpiSSHConfig()
}

func (c *connectorImpl) createServiceClients(config config.BmcsProperties, basePath string) error {

	authCSClient, err := c.authenticatedHttpsClient(coreServices, basePath,
		config.TransportConfig(coreServices))
	if err != nil {
		c.logger.Error(logTag, "Error connecting to service: %s. Reason: %v", coreServices, err)
		return err
	}
	cs := cclient.New(authCSClient, strfmt.Default)

	authIdentityClient, err := c.authenticatedHttpsClient(identityService, basePath,
		config.TransportConfig(identityService))
	if err != nil {
		c.logger.Error(logTag, "Error connecting to service %s. Reason: %v", identityService, err)
		return err
	}
	iam := iclient.New(authIdentityClient, strfmt.Default)

	c.coreService = cs
	c.iamService = iam
	return nil
}

func (c *connectorImpl) createCoreServiceClient() {

}

func (c *connectorImpl) createIdentityServiceClient() {

}

func (c *connectorImpl) authenticatedHttpsClient(host string, basePath string, config transport.Config) (*rclient.Runtime, error) {
	c.logger.Debug(logTag, "Creating Runtime Client")
	rt := rclient.New(host, basePath, []string{scheme})

	c.logger.Debug(logTag, "Creating authenticating transport to host %s", host)
	authC, err := transport.CreateAuthenticatedHTTPTarget(rt.Transport, config)

	rt.Transport = authC
	return rt, err
}
