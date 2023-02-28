/*
 * STACKIT DNS API
 *
 * This api provides dns
 *
 * API version: 1.0
 * Contact: stackit-dns@mail.schwarz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Zone.
type DomainZone struct {
	// access control list
	Acl string `json:"acl"`
	// contact email from soa record
	ContactEmail string `json:"contactEmail,omitempty"`
	// when zone creation finished
	CreationFinished string `json:"creationFinished"`
	// when zone creation started
	CreationStarted string `json:"creationStarted"`
	// default time to live
	DefaultTTL int32 `json:"defaultTTL"`
	// description of the zone
	Description string `json:"description"`
	// zone name
	DnsName string `json:"dnsName"`
	// Error shows error in case create/update/delete failed
	Error_ string `json:"error,omitempty"`
	// expire time
	ExpireTime int32 `json:"expireTime"`
	// zone id
	Id string `json:"id"`
	// if the zone is a reverse zone or not
	IsReverseZone bool `json:"isReverseZone,omitempty"`
	// user given name
	Name string `json:"name"`
	// negative caching
	NegativeCache int32 `json:"negativeCache"`
	// primary name server for secondary zone
	Primaries []string `json:"primaries,omitempty"`
	// primary name server. FQDN
	PrimaryNameServer string `json:"primaryNameServer"`
	// record count how many records are in the zone
	RecordCount int32 `json:"recordCount,omitempty"`
	// refresh time
	RefreshTime int32 `json:"refreshTime"`
	// retry time
	RetryTime int32 `json:"retryTime"`
	// serial number
	SerialNumber int32 `json:"serialNumber"`
	// zone state
	State string `json:"state"`
	// zone type
	Type_ string `json:"type"`
	// when zone update/deletion finished
	UpdateFinished string `json:"updateFinished"`
	// when zone update/deletion started
	UpdateStarted string `json:"updateStarted"`
	// visibility of the zone
	Visibility string `json:"visibility"`
}