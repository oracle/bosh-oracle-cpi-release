# BOSH Oracle CPI Release Repository
## Oracle Cloud Infrastructure Terminology (OCI)
  Consumers of this CPI are expected to be familiar with following OCI concepts
  * Tenancy 
  * Compartments 
  * Virtual Cloud Network(VCN) and Subnets
  * Regions and Availabilty Domains
  * Instance and Shapes
  * API signing key pair
  * Instance SSH key pair 
  
  See [OCI Concepts Documentation](https://docs.us-phoenix-1.oraclecloud.com/Content/GSG/Concepts/concepts.htm) to learn about these concepts
  
## Deploying BOSH Director 
   Refer to [Deploying BOSH Director Guide](docs/deploy_director.md)
      
## Development 
### Pre-requisites

Install [BOSH CLI v2](https://bosh.io/docs/cli-v2.html#install)

### Building a development release tarball 

#### Setup 
##### Clone this repo with the CPI submodule
```bash
$ git clone --recursive git@github.com:oracle/bosh-oracle-cpi-release.git
```

Note: CPI implementation source code lives in a different repository that is included as a git submodule under ./src of this 
repository 
##### Download Golang SDK into the local blobstore
````bash
$ ./populate-blobstore.sh
````
#### Build 
```bash
$ ./build-release.sh
```
build-release.sh will bump the dev release version and produce 

**bosh-oracle-cpi-0+dev.latest.tgz** 

under ./dev_releases/bosh-oracle-cpi directory

#### Clean up
```bash
$ bosh reset-release 
```
to remove all the versions of this release under ./dev_releases 


