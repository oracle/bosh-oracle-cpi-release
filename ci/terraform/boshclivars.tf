
resource "local_file"  bosh_cli_director_env_vars {
  depends_on = ["oci_identity_policy.bosh_policy"]

  filename = "./infra.yml"
  content = <<EOF
tenancy: ${var.oracle_tenancy_ocid}
user: ${oci_identity_user.bosh_user.id}
compartment: ${var.bosh_compartment["id"]}
apikey: |${format("\n   %s",join("\n   ", split("\n", file(var.bosh_api_private_key))))}
fingerprint: ${file(var.bosh_api_fingerprint)}
provisioned_username: ${var.bosh_ssh_username}
ssh_key: ${file(var.bosh_ssh_public_key)}
region: ${var.oracle_region}
ad1: ${oci_core_subnet.director_subnet_ad1.availability_domain}
vcn: "${lookup(data.oci_core_virtual_networks.VCNs.virtual_networks[0], "display_name")}"
subnet: ${oci_core_subnet.director_subnet_ad1.display_name}
internal_cidr: ${oci_core_subnet.director_subnet_ad1.cidr_block}
internal_gw: ${cidrhost(oci_core_subnet.director_subnet_ad1.cidr_block, 1)}
internal_ip: ${cidrhost(oci_core_subnet.director_subnet_ad1.cidr_block, 2)}
admin_password: admin
blobstore_agent_password: agent1
blobstore_director_password: director1
director_name: bosh-director
hm_password: hm1
mbus_bootstrap_password: mbus-secret
nats_password: nats-secret
postgres_password: postgres
EOF
}

resource "local_file"  bat_yml {
  depends_on = ["oci_identity_policy.bosh_policy"]

  filename = "./bat.yml"
  content = <<EOF
azs:
- name: z1
  cpi: oracle-cpi-ref
  cloud_properties:
    availability_domain: ((ad1))

vm_types:
- name: default
  cloud_properties:
    instance_shape: VM.Standard1.2
- name: vm4
  cloud_properties:
    instance_shape: VM.Standard1.4
- name: bm_default
  cloud_properties:
    instance_shape: BM.Standard1.36

disk_types:
- name: default
  disk_size: 51200
- name: large
  disk_size: 262144

networks:
- name: default
  type: manual
  subnets:
  - range: ((internal_cidr))
    gateway: ((internal_gw))
    az: z1
    dns: [8.8.8.8]
    reserved: [((internal_gw))/28]
    static: [${cidrhost(oci_core_subnet.director_subnet_ad1.cidr_block, 10)}-${cidrhost(oci_core_subnet.director_subnet_ad1.cidr_block, 128)}]
    cloud_properties:
      vcn: ((vcn))
      subnet_name: ((subnet))
compilation:
  workers: 5
  reuse_compilation_vms: true
  az: z1
  vm_type: default
  network: default
EOF
}
