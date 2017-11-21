# Authentication
variable "oracle_tenancy_ocid" {}
variable "oracle_user_ocid" {}
variable "oracle_fingerprint" {}
variable "oracle_private_key_path" {}

# Cloud services
variable oracle_region {
    default = "us-phoenix-1"
}


# Identity
variable "bosh_compartment" {
    type = "map"
}
variable "bosh_user_name" {
    default = "buildbot"
}
variable "bosh_group_name" {
    default = "buildbot"
}

variable "bosh_api_public_key" {
    default = "./keys/bosh-api-public-key.pem"
}

variable "bosh_api_private_key" {
    default = "./keys/bosh-api-private-key.pem"
}

variable "bosh_api_fingerprint" {
    default = "./keys/bosh-api-fingerprint"
}

variable "bosh_ssh_public_key" {
    default = "./keys/bosh-ssh.pub"
}
variable "bosh_ssh_private_key" {
    default = "./keys/bosh-ssh"
}
variable "bosh_ssh_username" {
    default = "ubuntu"
}

# Networking

variable "vpc_cidr" {
    default = "10.0.0.0/16"
}

variable "director_subnet_ad1_cidr" {
    default = "10.0.3.0/24"
}
