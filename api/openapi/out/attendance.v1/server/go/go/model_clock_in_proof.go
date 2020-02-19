/*
 * attendance
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ClockInProof struct {

	// A ID to a picture taken and uploaded to our servers
	PictureID string `json:"pictureID,omitempty"`

	// A ID generated by a device at the location, can be followed by coworkers coperating
	GeneratedID string `json:"generatedID,omitempty"`

	GPSID string `json:"GPSID,omitempty"`

	BiometricID string `json:"biometricID,omitempty"`
}
