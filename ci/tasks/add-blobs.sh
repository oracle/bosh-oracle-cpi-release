#!/usr/bin/env bash

function addGolangBlobToRelease() {
   local golang_distro=`basename /downloads/golang/go*.gz`
   local release_dir=$1

   echo "Adding blob golang/${golang_distro} to release ${release_dir}"
   bosh add-blob --dir ${release_dir} /downloads/golang/${golang_distro} golang/${golang_distro}

}