/*
 * Zplat API
 *
 * This is a zplat api-server api.
 *
 * API version: 1.0.0
 * Contact: saltbo@foxmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type ModelUser struct {
	Created string `json:"created,omitempty"`
	Email string `json:"email,omitempty"`
	Id int32 `json:"id,omitempty"`
	Role string `json:"role,omitempty"`
	Status string `json:"status,omitempty"`
	Ticket string `json:"ticket,omitempty"`
	Updated string `json:"updated,omitempty"`
	Username string `json:"username,omitempty"`
	// Global unique user ID
	Ux string `json:"ux,omitempty"`
}
