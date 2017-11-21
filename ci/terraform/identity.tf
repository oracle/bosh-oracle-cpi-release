resource "oci_identity_group" "bosh_group" {
    name = "${var.bosh_group_name}"
    description = "${var.bosh_group_name}"
}

resource "oci_identity_user" "bosh_user" {
    name = "${var.bosh_user_name}"
    description = "${var.bosh_user_name}"
}

resource "oci_identity_user_group_membership" "bosh_user_group_membership" {
    compartment_id = "${var.bosh_compartment["id"]}"
    user_id = "${oci_identity_user.bosh_user.id}"
    group_id = "${oci_identity_group.bosh_group.id}"
}

resource "oci_identity_api_key" "bosh_api_key" {
    user_id = "${oci_identity_user.bosh_user.id}"
    key_value = "${file(var.bosh_api_public_key)}"
}

resource "oci_identity_policy" "bosh_policy" {
    name = "bosh-policy"
    description = "bosh policies"
    compartment_id = "${var.bosh_compartment["id"]}"
    statements = [
        "allow group ${oci_identity_group.bosh_group.name} to manage instance-family in compartment ${var.bosh_compartment["name"]}",
        "allow group ${oci_identity_group.bosh_group.name} to manage volume-family in compartment ${var.bosh_compartment["name"]}",
        "allow group ${oci_identity_group.bosh_group.name} to manage object-family in compartment ${var.bosh_compartment["name"]}",
        "allow group ${oci_identity_group.bosh_group.name} to manage virtual-network-family in compartment ${var.bosh_compartment["name"]}",
        "allow group bosh-build-group to manage instance-family in compartment ${var.bosh_compartment["name"]}",
        "allow group bosh-build-group to manage object-family in compartment ${var.bosh_compartment["name"]}"
    ]
}
