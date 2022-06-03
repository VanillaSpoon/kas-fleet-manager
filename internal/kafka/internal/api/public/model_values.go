/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// Values struct for Values
type Values struct {
	Timestamp int64   `json:"timestamp,omitempty"`
	Value     float64 `json:"value"`
}
