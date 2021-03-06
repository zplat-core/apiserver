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

type ModelUserProfile struct {
	Avatar string `json:"avatar,omitempty"`
	Bio string `json:"bio,omitempty"`
	Company string `json:"company,omitempty"`
	Created string `json:"created,omitempty"`
	Id int32 `json:"id,omitempty"`
	Location string `json:"location,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Updated string `json:"updated,omitempty"`
	Url string `json:"url,omitempty"`
	Ux string `json:"ux,omitempty"`
}
