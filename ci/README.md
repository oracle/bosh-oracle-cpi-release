# BOSH Oracle CPI Release Concourse Pipeline

This document contains information for creating/running a [Concourse pipeline](https://pivotal.io/concourse) to produce Oracle's CPI for running Cloud Foundry on OCI. 

## Reference/Links

* [OCI Concepts](https://docs.us-phoenix-1.oraclecloud.com/Content/GSG/Concepts/concepts.htm)
* [Concourse CI](https://concourse.ci/introduction.html)
* [Fly CLI](https://concourse.ci/fly-cli.html)

## Prerequisites

* [OCI SDK/API Access](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/sdkconfig.htm)
* [Amazon S3 Compatibility API Key](https://docs.us-phoenix-1.oraclecloud.com/Content/Object/Tasks/s3compatibleapi.htm?Highlight=s3#APIsupport) 

## Pipeline Setup
 
### Login to concourse
```bash
$ fly -t nickname login -c http://concourse-host:port
```

### Create/update a pipeline
```bash
$ fly -t nickname set-pipeline --pipeline somename --config pipelines/cpi-release-pipeline.yml --load-vars-from private.yml
```

Note: More on the contents of `private.yml` below

#### private.yml
##### OCI parameters 
 
* `oracle-namespace`: Tenancy name  
* `oracle-tenancy`: Tenancy OCID  
* `oracle-region`: OCI Region name
* `oracle-s3-access-key-id`: OCID of the S3 Access Key
* `oracle-s3-secret-access-key`: Contents of the S3 Access Key
* `oracle-user`: User OCID
* `oracle-ad`: Availability domain name
* `oracle-fingerprint`: Fingerprint of the API key
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
* `oracle-fixture-env-bucket-name`: Name of the storage bucket used by terraform to create/delete OCI artifacts
* `stemcell-bucket`: Name of the storage bucket that will contain stemcells

##### BOSH/BATs deployment parameters
* `director-ad`: Name of the availability domain in which to install BOSH Director
* `director-compartment-name`: Name of the OCI compartment to use 
* `director-vcn`: Name of an existing VCN within the above compartment
* `director-subnet-cidr`: CIDR of the subnet to create for use by the BOSH director
* `bats-subnet1-cidr`: CIDR of BATs subnet1
* `bats-subnet2-cidr`: CIDR of BATs subnet2

