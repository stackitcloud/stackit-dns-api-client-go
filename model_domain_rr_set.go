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

// RRSet.
type DomainRrSet struct {
	// if the record set is active or not
	Active bool `json:"active,omitempty"`
	// comment
	Comment string `json:"comment,omitempty"`
	// when record set creation finished
	CreationFinished string `json:"creationFinished"`
	// when record set creation started
	CreationStarted string `json:"creationStarted"`
	// Error shows error in case create/update/delete failed
	Error_ string `json:"error,omitempty"`
	// rr set id
	Id string `json:"id"`
	// name of the record which should be a valid domain according to rfc1035 Section 2.3.4
	Name string `json:"name"`
	// records
	Records []DomainRecord `json:"records"`
	// record set state
	State string `json:"state"`
	// time to live
	Ttl int32 `json:"ttl"`
	// record set type
	Type_ string `json:"type"`
	// when record set update/deletion finished
	UpdateFinished string `json:"updateFinished"`
	// when record set update/deletion started
	UpdateStarted string `json:"updateStarted"`
}