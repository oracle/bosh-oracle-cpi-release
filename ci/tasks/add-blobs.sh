#!/usr/bin/env bash

echo "Adding blobs..."
golang_distro=`basename /downloads/golang/go*.gz`
bosh add-blob --dir ${release_dir} /downloads/golang/${golang_distro} golang/${golang_distro}