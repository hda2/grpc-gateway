/* 
 * A Bit of Everything
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * Contact: none@example.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package abe

// this message mimics string type, however when put in response rendered as a definition
type ExamplepbMimicObjectResponse struct {

	ResponseValueOne string `json:"response_value_one,omitempty"`

	ResponseValueTwo string `json:"response_value_two,omitempty"`

	ResponseEnum MimicObjectResponseMimicResponseEnum `json:"response_enum,omitempty"`
}
