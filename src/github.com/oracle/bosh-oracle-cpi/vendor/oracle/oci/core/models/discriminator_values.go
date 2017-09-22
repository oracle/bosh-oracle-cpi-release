package models

// DiscrimnatorTypeValues contains the mapping
// between a subtype name and the discriminator value used to
// identify that type when persisting or marshalling/unmarshalling instances
// of that type
var DiscriminatorTypeValues = map[string]string{

	"AttachIScsiVolumeDetails":                "iscsi",
	"IScsiVolumeAttachment":                   "iscsi",
	"DhcpDnsOption":                           "DomainNameServer",
	"DhcpSearchDomainOption":                  "SearchDomain",
	"ExportImageViaObjectStorageTupleDetails": "objectStorageTuple",
	"ExportImageViaObjectStorageUriDetails":   "objectStorageUri",
	"ImageSourceViaObjectStorageTupleDetails": "objectStorageTuple",
	"ImageSourceViaObjectStorageUriDetails":   "objectStorageUri",
}
