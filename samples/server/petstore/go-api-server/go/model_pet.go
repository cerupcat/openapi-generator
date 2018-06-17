/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

// Pet - A pet for sale in the pet store
type Pet struct {

	Id int64 `json:"id,omitempty"`

	Category *Category `json:"category,omitempty"`

	Name string `json:"name"`

	PhotoUrls []string `json:"photoUrls"`

	Tags []Tag `json:"tags,omitempty"`

	// pet status in the store
	Status string `json:"status,omitempty"`
}
