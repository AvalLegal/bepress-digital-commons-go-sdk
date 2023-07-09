/**
* Author: Bryson Meiling
* File: query.go
Get current results from your site that meet specific query parameters. Queries return up to 100
results by default. You can adjust this number from 1 to 1000 using the “limit” parameter, described
below. To retrieve all results, you may use the “export” endpoint
You can retrieve “query” results as often as you’d like, and the data is updated in real time as
changes and revisions are indexed in the Solr database.
*/

package lib

import (
	"net/http"
)

// represent optional url query paramters than can be added into the query
// q - the query string
// abstract - a string that is contained in the abstract
// author - a string that is contained in the author name
// publication_year - a year (YYYY) that is contained in the publication date
// parent_link - Retrieve all items in a publication context like a series or journal, using the full publication URL as a filter. For example, The format of a query that would return metadata for items in a collection is https://content-out.bepress.com/v2/{site_url}/query?parent_link=http://{site_url}/{publication}
// fields - If you would like to have a specific set of metadata fields returned, use the “fields” parameter. For example, this query for 'water' would return only the url, title, abstract, and autor of each hit. https://content-out.bepress.com/v2/{site_url}/query?q=water&fields=url,title,abstract,author
// select_fields - If you would like to return all available metadata fields, use select_fields=all. Keep in mind that `select_fields` and `fields` are mutually exclusive. If you use both, `fields` will be ignored.
// limit - Specifies how many results each page should contain; 100 results are returned by default. The limit can be set from 1 to 1000 results per page.
// start - Specifies from which result the page should start; the default value is 0.
type QueryParameters struct {
	Q               string
	Abstract        string
	Author          string
	PublicationYear string
	ParentLink      string
	Fields          string
	SelectFields    string
	Limit           string
	Start           string
}



// Query executes a query on the Digital Commons API with optional query parameters.
func (api *digitalCommonsAPI) Query(queryParams QueryParameters) ([]byte, error) {
	if err := validateQueryParameters(queryParams); err != nil {
		return nil, err
	}

	endpoint := "query"

	qp := map[string]string{
		"q":               queryParams.Q,
		"abstract":        queryParams.Abstract,
		"author":          queryParams.Author,
		"publication_year": queryParams.PublicationYear,
		"parent_link":     queryParams.ParentLink,
		"fields":          queryParams.Fields,
		"select_fields":   queryParams.SelectFields,
		"limit":           queryParams.Limit,
		"start":           queryParams.Start,
	}

	url := api.buildRequestURL(endpoint, qp)

	return api.makeRequest(url, http.MethodGet)
}
