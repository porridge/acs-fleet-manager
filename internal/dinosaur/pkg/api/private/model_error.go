/*
 * Red Hat Advanced Cluster Security Service Fleet Manager
 *
 * Red Hat Advanced Cluster Security (RHACS) Service Fleet Manager APIs that are used by internal services e.g fleetshard operators.
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// Error struct for Error
type Error struct {
	Id          string `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Href        string `json:"href,omitempty"`
	Code        string `json:"code,omitempty"`
	Reason      string `json:"reason,omitempty"`
	OperationId string `json:"operation_id,omitempty"`
}
