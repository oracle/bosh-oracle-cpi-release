#!/bin/bash
set -e 
LATEST_DEVTAR_NAME=bosh-oracle-cpi-0+dev.latest.tgz
MYDIR=`dirname $0`
RELEASE_PATH=$MYDIR/dev_releases/bosh-oracle-cpi
LATEST_TAR_PATH=$RELEASE_PATH/$LATEST_DEVTAR_NAME

function createReleaseV1() {

  bosh create release --force --with-tarball
  LAST_VERSION=`ls -1 -rt $RELEASE_PATH | tail -1`

  if [ -h $LATEST_TAR_PATH ]; then
    rm  $LATEST_TAR_PATH
  fi
  #Update latest release link

  ln -s $RELEASE_PATH/$LAST_VERSION $LATEST_TAR_PATH

}

function createReleaseV2() {

  if [ -h $LATEST_TAR_PATH ]; then
    rm  $LATEST_TAR_PATH
  fi

  bosh create-release --force --tarball=$LATEST_TAR_PATH
}

pushd $MYDIR
bosh_version=`bosh --version`
echo
if [[ $bosh_version == "BOSH 1".* ]]; then
   createReleaseV1
else
   createReleaseV2
fi
popd
