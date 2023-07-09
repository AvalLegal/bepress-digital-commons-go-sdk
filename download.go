/**
* Author: Bryson Meiling
* File: download.go
 */

package lib

import (
	"net/http"
)

// Download executes a download on the Digital Commons API with the provided export ID using the GET method.
func (api *digitalCommonsAPI) Download(exportID string) ([]byte, error) {
	endpoint := "download"
	url := api.buildRequestURL(endpoint, nil)
	url += "/" + exportID

	return api.makeRequest(url, http.MethodGet)
}