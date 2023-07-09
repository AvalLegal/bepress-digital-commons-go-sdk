/**
* Author: Bryson Meiling
* File: validateParameters.go
Used to validate some of the url parameters for user convenience and debugging
*/

package lib

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// validateQueryParameters validates the QueryParameters struct.
func validateQueryParameters(params QueryParameters) error {
	// Check if publication_year is a valid year
	if params.PublicationYear != "" {
		if _, err := strconv.Atoi(params.PublicationYear); err != nil {
			return errors.New("publication_year should be a valid year")
		}
	}

	// Check if select_fields and fields are mutually exclusive
	if params.SelectFields != "" && params.Fields != "" {
		return errors.New("select_fields and fields are mutually exclusive")
	}

	// Check if limit is a valid integer between 1 and 1000
	if params.Limit != "" {
		limit, err := strconv.Atoi(params.Limit)
		if err != nil || limit < 1 || limit > 1000 {
			return errors.New("limit should be an integer between 1 and 1000")
		}
	}

	// Check if start is a valid positive integer
	if params.Start != "" {
		start, err := strconv.Atoi(params.Start)
		if err != nil || start < 0 {
			return errors.New("start should be a positive integer")
		}
	}

	// Check if fields is a comma-separated string with no spaces
	if params.Fields != "" {
		fields := strings.Split(params.Fields, ",")
		regex := regexp.MustCompile(`^\s*$`)
		for _, field := range fields {
			if regex.MatchString(field) {
				return errors.New("fields should be a comma-separated string with no spaces")
			}
		}
	}

	return nil
}

// uses the validateQueryParameters function to validate the ExportParameters struct and additionally checks that `limit` and `start` are not set
func validateExportParameters(params QueryParameters) error {
	if err := validateQueryParameters(params); err != nil {
		return err
	}

	if params.Limit != "" {
		return errors.New("limit is not a valid parameter for export")
	}

	if params.Start != "" {
		return errors.New("start is not a valid parameter for export")
	}

	return nil
}


