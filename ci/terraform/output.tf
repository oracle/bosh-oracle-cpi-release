/*
output subnet_id {
  value = "${oci_core_subnet.director_subnet.id}"

}
output compartment_id {
   value = "${oci_core_subnet.director_subnet.compartment_id}"
}

output ad {
  value = "${oci_core_subnet.director_subnet.availability_domain}"
}

output subnet_name {
  value = "${oci_core_subnet.director_subnet.display_name}"
}

output subnet_gw {
  value = "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 1)}"
}

output first_ip {
   value = "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 2)}"
}
*/

output director_subnet {
  value = <<EOS
{
   "subnet_id" : "${oci_core_subnet.director_subnet.id}",
   "compartment_id" : "${oci_core_subnet.director_subnet.compartment_id}",
   "ad" : "${oci_core_subnet.director_subnet.availability_domain}",
   "subnet_name" : "${oci_core_subnet.director_subnet.display_name}",
   "subnet_cidr" : "${oci_core_subnet.director_subnet.cidr_block}",
   "subnet_gw" : "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 1)}",
   "director_ip" : "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 2)}"
}
EOS
}
