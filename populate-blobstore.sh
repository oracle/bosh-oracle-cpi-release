#!/bin/bash
set -e -x
#Script to download Go and add it to the 
#configured blobstore 

MYDIR=`dirname $0`
DOWNLOAD_DIR=$MYDIR/tmp
PACKAGE=golang
PACKAGE_VERSION=1.8.1
PLATFORMS="linux darwin"
create_dir=true
keep_download=false

function downloadPackage() {
  local blobname=`basename $1`
  pushd $1
  if [ -f $blobname ] ; then
    echo "Found $blobname in $2. Skipping download"
  else
    echo "Downloading $2"
    curl -L -O $2
  fi
  popd
}

function createDownloadDir() {
  mkdir -p $1
}

function removeDownloadDir() {
   rm -rf $1
}

function addGoDistributionToBlobstore() {
    local downloaddir=$1 platform=$2
    local downloadUrl=https://storage.googleapis.com/golang/go$PACKAGE_VERSION.$platform-amd64.tar.gz
    local blobname=`basename $downloadUrl`

    downloadPackage $downloaddir $downloadUrl

    bosh add-blob $downloaddir/$blobname $PACKAGE/$blobname

}

function downloadGoDistributions() {
   for platform in $PLATFORMS
   do
      addGoDistributionToBlobstore $1 $platform
   done
}

function parseOpts() {
   while getopts "d:k" opt; do
     case $opt in
       d)
         DOWNLOAD_DIR=$OPTARG
         create_dir=false
         if [ ! -d $DOWNLOAD_DIR ]; then
             echo "Invalid dir $DOWNLOAD_DIR" >&2
             exit 1
         fi
         ;;
       k)
         keep_download=true
         ;;
       \?)
         echo "Invalid option: -$OPTARG" >&2
         echo "Usage $0 [-d download_dir] [-k]"
         exit 1
         ;;
     esac
   done
}


parseOpts $@

if [ $create_dir == "true" ]; then 
  createDownloadDir $DOWNLOAD_DIR
fi

downloadGoDistributions $DOWNLOAD_DIR

if [ $create_dir == "true" -a $keep_download == "false" ] ; then
  removeDownloadDir $DOWNLOAD_DIR 
fi
