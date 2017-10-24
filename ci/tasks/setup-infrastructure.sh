#!/usr/bin/env bash
set -e

source bosh-cpi-src/ci/tasks/utils.sh

check_param oracle_tenancy
check_param oracle_user
check_param oracle_compartment_id
check_param oracle_compartment_name
check_param oracle_region
check_param oracle_ad
check_param oracle_fingerprint
check_param oracle_apikey
check_param oracle_ssh_public_key

infra_dir="$PWD/infra"

terraform_url="https://releases.hashicorp.com/terraform/0.10.6/terraform_0.10.6_linux_amd64.zip"
terraform_zip="/tmp/terraform.zip"

echo "Installing terraform..."
curl -Ls -o ${terraform_zip} ${terraform_url}
unzip -d /usr/local/bin ${terraform_zip}
chmod +x /usr/local/bin/terraform

echo "Installing oci provider for terraform..."
terraform_oci_provider_url="https://github.com/oracle/terraform-provider-oci/releases/download/v2.0.0/linux.tar.gz"
terraform_oci_provider_targz="/tmp/oci.tar.gz"
curl -Ls -o ${terraform_oci_provider_targz} ${terraform_oci_provider_url}

terraform_plugins_dir="$HOME/.terraform.d/plugins"
mkdir -p ${terraform_plugins_dir}
pushd ${terraform_plugins_dir}
  tar zxof ${terraform_oci_provider_targz}
popd

echo "Creating oci config..."
OCI_DIR="$HOME/.oci"
OCI_API_KEY="$OCI_DIR/oci_api_key.pem"
OCI_CONFIG="$OCI_DIR/config"

mkdir -p $OCI_DIR
cat > $OCI_API_KEY <<EOF
${oracle_apikey}
EOF

cat > $OCI_CONFIG <<EOF
[DEFAULT]
user=${oracle_user}
tenancy=${oracle_tenancy}
region=${oracle_region}
key_file=$OCI_API_KEY
fingerprint=${oracle_fingerprint}
EOF

chmod 600 $OCI_API_KEY


# Copy Terraform files
cp -r bosh-cpi-src/ci/terraform/* ${infra_dir}

keys_dir="${infra_dir}/keys"
mkdir -p ${keys_dir}
# Generate OCI API Key for bosh user.
openssl genrsa -out ${keys_dir}/bosh-api-private-key.pem 2048
chmod go-r ${keys_dir}/bosh-api-private-key.pem
openssl rsa -pubout -in ${keys_dir}/bosh-api-private-key.pem -out ${keys_dir}/bosh-api-public-key.pem
openssl rsa -pubout -outform DER -in ${keys_dir}/bosh-api-private-key.pem | openssl md5 -c | sed -e 's/(stdin)=\s*//g' > ${keys_dir}/bosh-api-fingerprint

# Generate SSL Certificates for Load Balancers.
#openssl genrsa -des3 -out ${keys_dir}/lb-ssl.key 2048
#openssl rsa -in ${keys_dir}/lb-ssl.key -out ${keys_dir}/lb-ssl.key
#openssl req -new -key ${keys_dir}/lb-ssl.key -out ${keys_dir}/lb-ssl.csr
#openssl x509 -req -days 365 -in ${keys_dir}/lb-ssl.csr -signkey ${keys_dir}/lb-ssl.key -out ${keys_dir}/lb-ssl.crt

echo "${oracle_ssh_public_key}" > ${keys_dir}/bosh-ssh.pub

# Init and run terraform
pushd ${infra_dir}
# Initialize vars whose values are passed in via the environment
cat > terraform.tfvars<<EOF
oracle_tenancy_ocid = "${oracle_tenancy}"
oracle_user_ocid = "${oracle_user}"
oracle_private_key_path = "$OCI_API_KEY"
oracle_fingerprint = "${oracle_fingerprint}"
oracle_region = "${oracle_region}"
bosh_compartment = {
  id = "${oracle_compartment_id}"
  name = "${oracle_compartment_name}"
}
EOF
terraform init
terraform plan
terraform apply
popd

# save the state/output of terraform for subsequent steps
tar zcf infra.tar.gz infra
oci os object put -ns cloudfoundry -bn infra --force --name infra.tar.gz --file infra.tar.gz

ls -al

#TODO bmcs CLI setup
oci os object list -ns cloudfoundry -bn infra
#oci compute instance list --compartment-id ${oracle_compartment_id}
