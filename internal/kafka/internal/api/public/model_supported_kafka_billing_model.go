/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.15.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// SupportedKafkaBillingModel Supported Kafka Billing Model
type SupportedKafkaBillingModel struct {
	// Identifier for the Kafka billing model
	Id string `json:"id"`
	// AMS resource to be used. Accepted values: ['rhosak']
	AmsResource string `json:"ams_resource"`
	// AMS product to be used. Accepted values: ['RHOSAK', 'RHOSAKTrial', 'RHOSAKEval', 'RHOSAKCC']
	AmsProduct string `json:"ams_product"`
	// List of AMS available billing models: Accepted values: ['marketplace', 'marketplace-rhm', 'marketplace-aws']
	AmsBillingModels []string `json:"ams_billing_models"`
}
