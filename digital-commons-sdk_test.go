/**
 * Author: Bryson Meiling
 * File: digital-commons-sdk_test.go
 */

package lib

import (
	"testing"
)
 
 func TestAPI(t *testing.T) {
	siteURL := "LawSchool"
	securtyToken := "secretKey123!"
	api := NewDigitalCommonsAPI(siteURL, securtyToken)

	// assert that the API was created correctly
	if api.baseURL != "https://content-out.bepress.com/v2" {
		t.Errorf("api.baseURL = %s; want https://content-out.bepress.com/v2", api.baseURL)
	}
	if api.siteURL != "LawSchool" {
		t.Errorf("api.siteURL = %s; want LawSchool", api.siteURL)
	}

 }

 func TestBuildRequestURL(t *testing.T) {
	siteURL := "LawSchool"
	securityToken := "secretKey123!"
	api := NewDigitalCommonsAPI(siteURL, securityToken)

	url := api.buildRequestURL("endpoint", nil)

	// asert that the url is correct
	if url != "https://content-out.bepress.com/v2/LawSchool/endpoint" {
		t.Errorf("url = %s; want https://content-out.bepress.com/v2/LawSchool/endpoint", url)
	}
 }

 func TestBuildRequestURLWithOneQueryParams(t *testing.T) {
	siteURL := "LawSchool"
	securityToken := "secretKey123!"
	api := NewDigitalCommonsAPI(siteURL, securityToken)

	qp := map[string]string{
		"q":               "my search query",
	}

	url := api.buildRequestURL("endpoint", qp)

	// asert that the url is correct
	if url != "https://content-out.bepress.com/v2/LawSchool/endpoint?q=my search query" {
		t.Errorf("url = %s; want https://content-out.bepress.com/v2/LawSchool/endpoint?q=my search query", url)
	}
 }

 func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

 func TestBuildRequestURLWithTwoQueryParams(t *testing.T) {
	siteURL := "LawSchool"
	securityToken := "secretKey123!"
	api := NewDigitalCommonsAPI(siteURL, securityToken)

	qp := map[string]string{
		"q": "a",
		"abstract": "b",
	}

	url := api.buildRequestURL("endpoint", qp)

	// asert that the url is correct
	var correctOptions []string
	correctOptions = append(correctOptions, "https://content-out.bepress.com/v2/LawSchool/endpoint?q=a&abstract=b")
	correctOptions = append(correctOptions, "https://content-out.bepress.com/v2/LawSchool/endpoint?abstract=b&q=a")

	// if url not in CorrectOptions
	if !contains(correctOptions, url) {
		t.Errorf("url = %s; want https://content-out.bepress.com/v2/LawSchool/endpoint?q=a&abstract=b or https://content-out.bepress.com/v2/LawSchool/endpoint?abstract=b&q=a", url)

	}
 }
 