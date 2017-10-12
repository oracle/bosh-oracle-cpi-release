# Oracle CPI and Cloud Config

This document describes the required Oracle cloud specific properties in [Cpi](https://bosh.io/docs/cpi-config.html) and [Cloud](https://bosh.io/docs/cloud-config.html) config for deploying a BOSH release to Oracle Cloud Infrastructure. 
  
## Cpi Config

#### cpis block

**type**:  CPI binary name. Director appends the string *"_cpi"* to this when looking for the cpi excutable so this must be set to *'oracle'*

**properties**:

  Name    | Description
  -----   | ----------
  tenancy | OCID of the tenancy to use  
  compartment | OCID of the compartment in which the CPI will create the requested resources
  user | OCID of the user in that tenancy. The user must have appropriate permissions to create resources in the above compartment
  apikey | Contents of the API signing private key file
  fingerprint |  Fingerprint of the API signing public key uploaded to OCI
  region | Name of the region to use
   
##### Example:

```yaml
cpis:
- name: oracle-cpi-ref
  type: oracle  #CPI binary name. Director appends _cpi to this name 
  properties:
    tenancy: ((tenancy))
    compartment: ((compartment))
    apikey: ((apikey))
    user: ((user))
    fingerprint: ((fingerprint))
    region: ((region))

```

## Cloud Config

#### azs block

   BOSH availability zones are mapped to OCI availability domains.
   
**cloud_properties:**


  Name | Description
  ---- | -----------
  availability_domain | Name of the availability_domain in the region configured in CPI config
    
#### vm_types block

Virtual machines or bare metal instances of varying cpu and memory capacity can be created in OCI. Capacity configurations are pre-defined as a Shape.

**cloud_properties:**


  Name | Description
  ---- | -----------
  instance_shape | Name of the shape to use

#### disk_types block
  This block doesn't require any Oracle cloud specific properties.  
  
  However disk_size must be set to 50GiB or a higher multiple. Volumes smaller than 50GiB are currently not supported in OCI.
   
#### networks block
   A BOSH network and its subnets are mapped to a VCN and its subnets respectively.  Currently only static (*type: manual*) and dynamic (*type: dynamic*) networks are supported i.e.,  (*type: vip*) isn't supported yet 
   
**cloud_properties:**

  Name | Description
  ---- | -----------
  vcn  | Name of the Virtual Cloud Network (VCN)
  subnet | Name of the subnet in the above vcn


##### Example:
```yaml
azs:
- name: z1
  cpi: oracle-cpi-ref
  cloud_properties:
    availability_domain: ((ad1))
- name: z2
  cpi: oracle-cpi-ref
  cloud_properties:
    availability_domain: ((ad2))
- name: z3
  cpi: oracle-cpi-ref
  cloud_properties:
    availability_domain: ((ad3))

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
  - range: ((internal_cidr_n1))
    gateway: ((internal_gw_n1))
    az: z1
    dns: [8.8.8.8]
    reserved: [((internal_gw_n1))/28]
    static: [10.0.1.32-10.0.1.128]
    cloud_properties:
      vcn: ((vcn))
      subnet_name: ((subnet1))
  - range: ((internal_cidr_n2))
    gateway: ((internal_gw_n2))
    az: z2
    dns: [8.8.8.8]
    reserved: [((internal_gw_n2))/28]
    static: [10.0.2.32-10.0.2.128]
    cloud_properties:
      vcn: ((vcn))
      subnet_name: ((subnet2))

  - range: ((internal_cidr_n3))
    gateway: ((internal_gw_n3))
    az: z3
    dns: [8.8.8.8]
    reserved: [((internal_gw_n3))/28]
    static: [10.0.3.32-10.0.3.128]
    cloud_properties:
      vcn: ((vcn))
      subnet_name: ((subnet3))

compilation:
  workers: 5
  reuse_compilation_vms: true
  az: z1
  vm_type: default
  network: default 

```
