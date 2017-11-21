
resource "oci_core_security_list" "director_all" {
    compartment_id        = "${var.bosh_compartment["id"]}"
    display_name          = "director_all"
    vcn_id                = "${lookup(data.oci_core_virtual_networks.VCNs.virtual_networks[0], "id")}"
    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        protocol = "6"
        source = "${var.director_subnet_ad1_cidr}"
    },
    {
        protocol = "1"
        source   = "${var.vpc_cidr}"
    },
    {
        tcp_options {
            "max" = 22
            "min" = 22
        }
        protocol = "6"
        source   = "${var.vpc_cidr}"
    },
    {
        tcp_options {
            "max" = 4222
            "min" = 4222
        }
        protocol = "6"
        source   = "${var.vpc_cidr}"
    },
    {
        tcp_options {
            "max" = 6868
            "min" = 6868
        }
        protocol = "6"
        source   = "${var.vpc_cidr}"
    },
    {
        tcp_options {
            "max" = 8443
            "min" = 8443
        }
        protocol = "6"
        source   = "${var.vpc_cidr}"
    },
    {
        tcp_options {
            "max" = 25250
            "min" = 25250
        }
        protocol = "6"
        source   = "${var.vpc_cidr}"
    },
    {
        tcp_options {
            "max" = 25555
            "min" = 25555
        }
        protocol = "6"
        source   = "${var.vpc_cidr}"
    },
    {
        tcp_options {
            "max" = 25777
            "min" = 25777
        }
        protocol = "6"
        source   = "${var.vpc_cidr}"
    }]
}

resource "oci_core_subnet" "director_subnet_ad1" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0], "name")}"
  cidr_block          = "${var.director_subnet_ad1_cidr}"
  display_name        = "director_subnet_ad1"
  dhcp_options_id     = "${lookup(data.oci_core_virtual_networks.VCNs.virtual_networks[0], "default_dhcp_options_id")}"
  dns_label           = "cfdirad1"
  compartment_id      = "${var.bosh_compartment["id"]}"
  vcn_id              = "${lookup(data.oci_core_virtual_networks.VCNs.virtual_networks[0], "id")}"
  route_table_id      = "${lookup(data.oci_core_virtual_networks.VCNs.virtual_networks[0], "default_route_table_id")}"
  security_list_ids   = ["${oci_core_security_list.director_all.id}"]
  prohibit_public_ip_on_vnic = false # https://jira.aka.lgl.grungy.us/browse/CF-229
}
