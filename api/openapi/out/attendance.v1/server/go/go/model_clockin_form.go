/*
 * attendance
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type ClockinForm struct {

	BusinessID string `json:"businessID"`

	ClockinTime time.Time `json:"clockinTime"`

	ClockoutTime time.Time `json:"clockoutTime"`

	ClockinProof []ClockInProof `json:"clockinProof,omitempty"`
}