# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.oracle_tenancy_ocid}"
}

data "oci_core_virtual_networks" "VCNs" {
    compartment_id = "${var.bosh_compartment["id"]}"
    limit = 1
}
