Repository for building BOSH Oracle CPI Release.
#### Pre-requisites

Install [BOSH CLI v2](https://bosh.io/docs/cli-v2.html#install)

#### Building the development release tarball 

```bash
$ build_release.sh
```
build_release.sh will produce **bosh-oracle-cpi-0+dev.latest.tgz** in the current directory

#### Clean up
```bash
$ bosh reset-release 
```
to remove all the versions of this release under dev_releases 


---

Note: CPI implementation source code lives in a different repository that is included as a git submodule under ./src of this 
repository 
