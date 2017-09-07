Deploy BOSH Director
====================

Pre-requisites
==============
- BOSH cli v2
- VCN, Subnet, and bastion instance in the compartment where  BOSH director should be deployed
- An API signing key uploaded to your account (using console or sdk)

- A variables file $HOME/director-env-vars.yml containing the values for the environment above and passwords to use for
BOSH director components


Create BOSH environment(deploy director)
==============================
- Run create-env.sh

Verify environment accessible
==================
- set-boshenv.sh <my-env-alias>
- bosh -e <my-env-alias>

Delete BOSH environment
==============================
- Run delete-env.sh
