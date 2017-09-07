## BOSH Oracle CPI Release Repository

### Pre-requisites

Install [BOSH CLI v2](https://bosh.io/docs/cli-v2.html#install)

### Building the development release tarball 

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


---

