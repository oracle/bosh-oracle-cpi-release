#!/bin/bash

set -e

DOCKER_IMAGE=${DOCKER_IMAGE:-dmutreja/oci-terraform-resource}
DOCKER_IMAGE_VERSION=${DOCKER_IMAGE_VERSION:-latest}

docker login -u $DOCKER_USER -p $DOCKER_PWD


if [ ! -d oci-provider ]; then
 echo "Downloading oci-provider..."
 mkdir oci-provider
 curl -sSL https://github.com/oracle/terraform-provider-oci/releases/download/v2.0.4/linux.tar.gz | tar -xzC oci-provider
fi

if [ ! -d null-provider ]; then 
 echo "Downloading null-provider..."
 mkdir -p null-provider/linux_amd64
 curl -sSL -O https://releases.hashicorp.com/terraform-provider-null/1.0.0/terraform-provider-null_1.0.0_linux_amd64.zip 
 unzip -d null-provider/linux_amd64 terraform-provider-null_1.0.0_linux_amd64.zip  
fi


echo "Building docker image..."
docker build -t $DOCKER_IMAGE .

echo "Pushing docker image to '$DOCKER_IMAGE'..."
docker push $DOCKER_IMAGE
