/*
 * message
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMessageAll - Your GET endpoint
func GetMessageAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetMessageMessageID - Your GET endpoint
func GetMessageMessageID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostMessageBusinessIDForm - 
func PostMessageBusinessIDForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}