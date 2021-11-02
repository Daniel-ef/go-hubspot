/*
 * CMS Source Code
 *
 * Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package source_code

import (
	"os"
)

type ValidatePathBody struct {
	// The file to validate.
	File **os.File `json:"file,omitempty"`
}
