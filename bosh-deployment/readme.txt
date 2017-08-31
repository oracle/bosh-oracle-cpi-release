Deploy BOSH Director
====================

Pre-requisites
==============
- Download bosh cli v2
- VCN, Subnet, and bastion instance in the cpmpartment where  BOSH director should be deployed
- Create an API signing key and upload it to your account (using console or sdk)

- Fill my-bmc-vars.yml based on the environment created above

Create BOSH environment(deploy director)
==============================
- Run create_env.sh

Verify environment accessible
==================
- set_boshenv.sh <my-env-alias>
- bosh -e <my-env-alias>

Delete BOSH environment
==============================
- Run delete_env.sh
