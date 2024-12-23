/*
 * STACKIT DNS API
 *
 * This api provides dns
 *
 * API version: 1.0
 * Contact: dns@stackit.cloud
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Record.
type DomainRecord struct {
	// content of the record
	Content string `json:"content"`
	// rr set id
	Id string `json:"id"`
}
