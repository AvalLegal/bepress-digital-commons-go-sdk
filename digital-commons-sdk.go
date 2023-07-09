/**
* Author: Bryson Meiling
* File: digital-commons-sdk.go
The Digital Commons API client for Go. Contains the shared logic for making requests to the Digital Commons API which is used by the 4 endpoints files: query.go, download.go, export.go, and fields.go.
*/

package lib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DigitalCommonsAPI represents the Digital Commons API client.
type digitalCommonsAPI struct {
	baseURL 	 string
	siteURL       string
	securityToken string
	httpClient    *http.Client
}

// NewDigitalCommonsAPI creates a new instance of the DigitalCommonsAPI client.
func NewDigitalCommonsAPI(siteURL string, securityToken string) *digitalCommonsAPI {
	httpClient := &http.Client{}
	const baseURL = "https://content-out.bepress.com/v2"
	return &digitalCommonsAPI{
		baseURL:       baseURL,
		siteURL:       siteURL,
		securityToken: securityToken,
		httpClient:    httpClient,
	}
}

// buildRequestURL constructs the API request URL with the given endpoint and query parameters.
func (api *digitalCommonsAPI) buildRequestURL(endpoint string, queryParams map[string]string) string {
	url := fmt.Sprintf("%s/%s/%s", api.baseURL, api.siteURL, endpoint)
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1] // Remove the trailing '&'
	}
	return url
}

// makeRequest makes a request to the Digital Commons API with the provided URL, headers, and HTTP method.
func (api *digitalCommonsAPI) makeRequest(url string, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// set the Authorization header
	req.Header.Set("Authorization", api.securityToken)

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and return the response body.
	// Handle any error cases based on the response status code.
	switch resp.StatusCode {
	case http.StatusOK:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized request")
	case http.StatusNotFound:
		return nil, errors.New("resource not found")
	case http.StatusInternalServerError:
		return nil, errors.New("internal server error")
	default:
		return nil, errors.New("unknown error")
	}

}







