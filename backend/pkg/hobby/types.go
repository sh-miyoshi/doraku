/*
 * Doraku
 *
 * Doraku Backend API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hobby

// Hobby struct defines info of hobby
type Hobby struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NameEN      string `json:"name_EN"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Group       int32  `json:"group"`
}
