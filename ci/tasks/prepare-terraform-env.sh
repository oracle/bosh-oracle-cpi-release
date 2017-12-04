#!/usr/bin/env bash

set -e

pwd=`pwd`
env_dir=${pwd}/terraform-env
api_key_path=terraform-env/oci_api_key.pem
vars_file=${env_dir}/oci.vars

mkdir -p $env_dir

echo "Creating terraform variables file..."

cat > ${api_key_path} <<EOF
${oracle_apikey}
EOF
chmod 600 ${api_key_path}

cat > ${vars_file} <<EOF
oracle_user_ocid: ${oracle_user}
oracle_tenancy_ocid: ${oracle_tenancy}
oracle_region: ${oracle_region}
oracle_private_key_path: ${api_key_path}
oracle_fingerprint: ${oracle_fingerprint}
director_compartment_name: ${director_compartment_name}
director_vcn: ${director_vcn}
director_subnet_cidr: ${director_subnet_cidr}
director_ad: ${director_ad}
bats_subnet1_cidr: ${bats_subnet1_cidr}
bats_subnet2_cidr: ${bats_subnet2_cidr}
EOF

echo "Done. Created: " ${vars_file}