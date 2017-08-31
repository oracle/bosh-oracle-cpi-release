    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***

# BOSH Oracle Bare Metal CPI

# EXPERIMENTAL

This is an external [BOSH CPI](http://bosh.io/docs/bosh-components.html#cpi) for [Oracle Bare Metal Cloud](https://cloud.oracle.com/en_US/bare-metal)

## Usage

### Installation
The CPI implementation contained in this repository is deployed as part of  [BOSH Oracle CPI release](https://gitlab-odx.oracle.com/cloudfoundry/bosh-oracle-cpi-release). 

### Development

The source for this CPI is not intended to be deployed except as a BOSH deployment. 
To build or install the CPI locally for development or test purposes, you can symlink the repository into your Go workspace.

From the root of the `bosh-oracle-cpi-release` repository:

```
ln -s $(pwd)/src/github.com/oracle/bosh-oracle-cpi $GOPATH/src/github.com/oracle/bosh-oracle-cpi
```

### Building the CPI locally

From $GOPATH/src/oracle/bosh/cpi

```
make build
```
This will build the cpi executable as $GOPATH/src/oracle/bosh/cpi/out/cpi

#### Running development tests

$GOPATH/src/oracle/bosh/cpi/Makefile has two targets to run development tests

1. test

    ``make test``

    Runs all unit tests. Unit tests use fakes and don't interact with the Bare Metal Cloud.


2. bmctest

    ``make bmctest``

    Runs tests in bmc/test package. The tests in this package drive the layer responsible for bulk of the work in the CPI. E.g. creating/deleting  a vm insatnce  or a persistent block volume
    in the Oracle Bare Metal Cloud. Configuration values required for running this target are specified in the format used by various [BMC SDKs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/sdkconfig.htm).

    The default configuration file name and location is ~/.oraclebmc/config.  The configuration file can contain multiple profiles. Tests run by this target
    use the [CPITEST] profile by default.

    The default file location and the profile name can be changed by
    setting the CPITEST_CONFIG and CPITEST_PROFILE environment variables respectively.

    List of configuration entries required

    | Entry         | Description
    | ------------- |-------------
    | tenancy       | OCID of the tenancy in which CPI will create(or delete) the requested resources
    | user          | OCID of the user in that tenancy 
    | region        | Region to use
    | key_file      | API key file used to sign the requests 
    | fingerprint   | Fingerprint of the API key uploaded to BMC
    | compartment   | OCID of the compartment in which the resources will be created
    | ad            | Name of the Avaialibilty Domain in the region
    | vcn           | Name of the VCN to use
    | subnet        | Name of the subnet in that VCN to use
    | cpiUser       | Name of the provisioned user in the image. Used by CPI ssh to attach iscsi block volumes
    | cpiPrivateKeyPath | Path to locally provisioned ssh private key
    | cpiPublicKeyPath | Path to locally provisioned ssh public key
    | userPublicKeyPath | Optional public key to install on newly created instance
    | stemcellImage | OCID of the stemcell image for creating new instances



#### Running the CPI directly

##### Configuration

Create a cpi.json configuration file:

##### Run
Run it using the previously created json
```
$ echo "{\"method\": \"method_name\", \"arguments\": []}" | cpi -configFile="/path/to/configuration_file.json"
```
### Installing BOSH using this CPI

See [Deploying BOSH Director](https://confluence.aka.lgl.grungy.us/pages/viewpage.action?pageId=14285397)

## Features

TODO: Document the various cloud properties available in the deployment manifest

#### BOSH Network options

#### BOSH Resource pool options

#### BOSH Persistent Disks options

#### Deployment Manifest Example


