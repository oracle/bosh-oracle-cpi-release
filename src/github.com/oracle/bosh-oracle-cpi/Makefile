# Uncomment here or export (in the launch shell) the following variables to control test configuration
#CPITEST_CONFIG=/path/to/my/oci/config ini. Default is ~/.oci/config
#CPITEST_PROFILE=section inside CPITEST_CONFIG file. Default is CPITEST

default: build

# Builds bosh-cpi for linux-amd64
build:
	go build -o out/cpi github.com/oracle/bosh-oracle-cpi/main

# Build cross-platform binaries
build-all:
	gox -output="out/cpi_{{.OS}}_{{.Arch}}" github.com/oracle/bosh-oracle-cpi/main 

# Prepration for tests
get-deps:
	# Go lint tool
	go get github.com/golang/lint/golint

	# Simplify cross-compiling
	go get github.com/mitchellh/gox

	# Ginkgo and omega test tools
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

# Cleans up directory and source code with gofmt
clean:
	go clean ./...

# Run gofmt on all code
fmt:
	gofmt -l -w .

# Run linter with non-stric checking
lint:
	@echo ls -d */ | grep -v vendor | xargs -L 1 golint
	ls -d */ | grep -v vendor | xargs -L 1 golint

# Vet code
vet:
	go tool vet $$(ls -d */ | grep -v vendor)

# Runs all the unit tests 
test: get-deps clean fmt lint vet build
	ginkgo -r -race -skipPackage=oci/test .

# Runs a single unit test spec specified in TEST_SPEC 
single-test: get-deps  
	ginkgo -r -race -skipPackage=oci/test -focus=$(TEST_SPEC) $(GINKGO_ARGS) .

# Runs OCI integration tests
ocitest: 
	ginkgo -slowSpecThreshold=500 -progress -nodes=3 -randomizeAllSpecs -randomizeSuites $(GINKGO_ARGS) -v oci/test

# Runs a subset of OCI integration tests
ocitest-subset:
	go test -v  ./oci/test/... -run $(TEST_SPEC)

# Runs the integration tests from Concourse
# Set TEST_SPEC to control which tests are run
#. E.g. to skip long running stemcell tests set 
# TEST_SPEC="Test_[^StemcellOps]"
testintci: 
	go test -v  ./oci/test/... -run $(TEST_SPEC) -timeout 60m
