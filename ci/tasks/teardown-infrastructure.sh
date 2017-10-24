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

# get the terraform artifacts that were created previously
oci os object get -ns cloudfoundry -bn infra --name infra.tar.gz --file infra.tar.gz
tar zxof infra.tar.gz

echo "Tearing down oracle cloud infrastructure..."
set +e

# Init and run terraform to undo previous changes
infra_dir="$PWD/infra"
pushd ${infra_dir}
terraform init
terraform plan -destroy
terraform destroy --force
popd
set -e
