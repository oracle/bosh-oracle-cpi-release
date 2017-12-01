#!/usr/bin/env bash

set -e

pwd=`pwd`

#Inputs
export TERRAFORM_OUTPUT=${pwd}/terraform-out/metadata
templates_path=${pwd}/cpi-release-src/ci/templates
keys=${pwd}/fixture-ssh-keys


#Output
output_dir=${pwd}/oci-config
apikey_file_name=ocitest_api_key.pem

key_path=${output_dir}/${apikey_file_name}
cat > ${key_path} <<EOF
${oracle_apikey}
EOF
chmod 600 ${key_path}
# oracle_key_path relative to
# the config file
export oracle_apikey_path="./${apikey_file_name}"

cp -pr ${keys}/* ${output_dir}/

erb -T '-' -r json ${templates_path}/ocitest-ini.erb >  ${output_dir}/config

export userPublicKeyPath=${output_dir}/userkeys/id_rsa.pub
erb -T '-' -r json ${templates_path}/create-env-vars.erb >  ${output_dir}/director-env-vars.yml

erb -T '-' -r json ${templates_path}/bat.yml.erb >  ${output_dir}/bat.yml