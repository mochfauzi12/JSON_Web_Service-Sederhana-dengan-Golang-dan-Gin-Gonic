package service_test

import (
	"Web_Service/internal/infrastructure/service"
	"testing"
)

func TestGetRandomUser_Error(t *testing.T) {
	_, err := service.GetRandomUser()
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseJSON_Error(t *testing.T) {
	invalidJSON := []byte(`{"invalid": "json"}`)
	_, err := service.ParseJSON(invalidJSON)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseJSON_NoResults(t *testing.T) {
	noResultsJSON := []byte(`{"results": []}`)
	_, err := service.ParseJSON(noResultsJSON)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != "NO RESULTS FOUND FOR RANDOM USER" {
		t.Errorf("Expected error message: 'NO RESULTS FOUND FOR RANDOM USER', but got: %s", err.Error())
	}
}
