#!/usr/bin/env bash

set -e

# Create OCI config
echo "Creating oci config..."
OCI_DIR="$HOME/.oci"
OCI_API_KEY="$OCI_DIR/oci_api_key.pem"
OCI_CONFIG="$OCI_DIR/config"

mkdir -p ${OCI_DIR}
cat > ${OCI_API_KEY} <<EOF
${apikey}
EOF

cat > $OCI_CONFIG <<EOF
[DEFAULT]
user=${user}
tenancy=${tenancy}
region=${region}
key_file=$OCI_API_KEY
fingerprint=${fingerprint}
EOF

chmod 600 ${OCI_API_KEY}
chmod 600 ${OCI_CONFIG}

# Download infra
oci os object get -ns cloudfoundry -bn ${bucket} --name infra.tar.gz --file infra-out/infra.tar.gz
