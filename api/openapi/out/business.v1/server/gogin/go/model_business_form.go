/*
 * business_service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BusinessForm struct {

	Name string `json:"name"`

	Location string `json:"location"`

	// A array of users who can modify this business
	UserIDs []string `json:"userIDs"`
}