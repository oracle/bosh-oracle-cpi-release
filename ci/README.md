# BOSH Oracle CPI Release Concourse Pipeline

This document contains information for creating/running a [Concourse pipeline](https://pivotal.io/concourse) to produce Oracle's CPI for running Cloud Foundry on OCI. 

## Reference/Links

* [OCI Concepts](https://docs.us-phoenix-1.oraclecloud.com/Content/GSG/Concepts/concepts.htm)
* [Concourse CI](https://concourse.ci/introduction.html)
* [Fly CLI](https://concourse.ci/fly-cli.html)

## Prerequisites

* [OCI SDK/API Access](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/sdkconfig.htm)
* [Amazon S3 Compatibility API Key](https://docs.us-phoenix-1.oraclecloud.com/Content/Object/Tasks/s3compatibleapi.htm?Highlight=s3#APIsupport) 

To use these instructions you need to setup the following on Oracle's Public Cloud:
- An Oracle Cloud Infrastructure account
- A user created in that account, in a group with a policy that grants the desired permissions. This can be a user for yourself, or another person/system that needs to call the API. Click [here](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#Other) for more information on this. 
- A key pair used for signing API requests, with the public key uploaded to Oracle. Only the user calling the API should be in possession of the private key. Click [here](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#How) for instructions on how to generate and upload a key pair.
- An OCI [Compartment](https://docs.us-phoenix-1.oraclecloud.com/Content/GSG/Tasks/choosingcompartments.htm?Highlight=compartment) in which the pipeline runs
- A [Virtual Cloud Network (VCN)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/overview.htm)
- [Create](http://www.oracle.com/webfolder/technetwork/tutorials/obe/cloud/ocis/object-storage/object-storage.html) the following object storage buckets
  * `cpi-dev-release-bucket`: Name of the storage bucket that will contain dev/candidate releases
  * `cpi-final-release-bucket`: Name of the storage bucket that will contain final releases
  * `version-semver-bucket-name`: Name of the storage bucket used to maintan dev/final release version numbers
  * `oracle-fixture-env-bucket-name`: Name of the storage bucket to store terraform state and ssh keys required for running tests
  * `stemcell-bucket`: Name of the storage bucket that contains stemcells
- Create an [Amazon S3 Compatibility API key](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingcredentials.htm#To4)
- A user with read access to the GitHub repositories (or a fork) [https://github.com/oracle/bosh-oracle-cpi-release] and [https://github.com/oracle/bosh-oracle-cpi]

## Pipeline Setup
 
### Login to concourse
```bash
$ fly -t nickname login -c http://concourse-host:port
```

### Create/update a pipeline
```bash
$ fly -t nickname set-pipeline --pipeline somename --config pipelines/cpi-release-pipeline.yml --load-vars-from env-vars.yml
```

Note: More on the contents of `env-vars.yml` below

#### Create env-vars.yml

Create a YAML file containing the values specific to the OCI environment in which concourse/BATS etc run. The various required fields are described below. 

##### OCI parameters 
 
* `oracle-namespace`: Tenancy name  
* `oracle-tenancy`: [Tenancy OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#Other)  
* `oracle-region`: [OCI Region name](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm)
* `oracle-s3-access-key-id`: OCID of the S3 Access Key
* `oracle-s3-secret-access-key`: Contents of the S3 Access Key
* `oracle-user`: [User OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#Other)
* `oracle-ad`: Availability domain name
* `oracle-fingerprint`: [Fingerprint of the API key](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#four)
* `oracle-apikey`: The contents of the API key

##### Source code parameters 
* `github-org`: The name of the github user/organization hosting the source code
* `github-user`: Github username
* `github-password`: Github password (for commit access)
* `cpi-release-branch`: Branch name of the CPI release repository
* `cpi-branch`: Branch name of the CPI repository
           
##### OCI object storage parameters 
* `cpi-dev-release-bucket`: Name of the storage bucket that will contain dev/candidate releases
* `cpi-final-release-bucket`: Name of the storage bucket that will contain final releases
* `version-semver-bucket-name`: Name of the storage bucket used to maintan dev/final release version numbers
* `oracle-fixture-env-bucket-name`: Name of the storage bucket to store terraform state and ssh keys required for running tests
* `stemcell-bucket`: Name of the storage bucket that contains stemcells

##### BOSH/BATs deployment parameters
* `director-ad`: Name of the [availability domain](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm) in which to install BOSH Director
* `director-compartment-name`: Name of the OCI compartment to use 
* `director-vcn`: Name of an existing VCN within the above compartment
* `director-subnet-cidr`: CIDR of the subnet to create for use by the BOSH director
* `bats-subnet1-cidr`: CIDR of BATs subnet1
* `bats-subnet2-cidr`: CIDR of BATs subnet2

