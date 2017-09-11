#!/bin/bash

if [ $# -ne 1 ]; then 
  echo "Usage $0 <target-env-alias>"
  exit 1
fi
bosh alias-env $1 -e $(bosh int $HOME/director-env-vars.yml --path /internal_ip) --ca-cert <(bosh int creds.yml --path /director_ssl/ca)
