/**
* Author: Bryson Meiling
* File: export.go
 */

package lib

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ExportID represents the export ID returned by the PUT request to the export endpoint.
type ExportID struct {
	ExportId string `json:"ExportId"`
}

// Export executes an export on the Digital Commons API with optional query parameters using the PUT method.
func (api *digitalCommonsAPI) Export(queryParams QueryParameters) (*ExportID, error) {
	if err := validateExportParameters(queryParams); err != nil {
		return nil, err
	}

	endpoint := "export"

	qp := map[string]string{
		"q":               queryParams.Q,
		"abstract":        queryParams.Abstract,
		"author":          queryParams.Author,
		"publication_year": queryParams.PublicationYear,
		"parent_link":     queryParams.ParentLink,
		"fields":          queryParams.Fields,
		"select_fields":   queryParams.SelectFields,
	}

	url := api.buildRequestURL(endpoint, qp)

	resp, err := api.makeRequest(url, http.MethodPut)
	if err != nil {
		return nil, err
	}

	// Parse the response JSON
	var jsonResponse map[string]string
	err = json.Unmarshal(resp, &jsonResponse)
	if err != nil {
		return nil, err
	}

	// Check if the response contains the expected JSON object structure
	exportID, ok := jsonResponse["ExportId"]
	if !ok || exportID == "" {
		return nil, errors.New("invalid export response")
	}

	return &ExportID{exportID}, nil
}