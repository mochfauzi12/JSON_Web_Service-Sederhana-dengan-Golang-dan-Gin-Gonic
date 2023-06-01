package service_test

import (
	"Web_Service/domain/entity"
	"Web_Service/internal/infrastructure/service"
	"testing"
)

func TestGetRandomUser_Success(t *testing.T) {
	body, err := service.GetRandomUser()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(body) == 0 {
		t.Error("Expected non-empty response body, but got empty")
	}
}

func TestParseJSON_Success(t *testing.T) {
	validJSON := []byte(`{
		"results": [
			{
				"gender": "male",
				"name": {
					"title": "Mr",
					"first": "John",
					"last": "Doe"
				},
				"location": {
					"street": "123 Main St",
					"city": "New York"
				},
				"picture": {
					"large": "https://example.com/profile.jpg"
				}
			}
		]
	}`)

	response, err := service.ParseJSON(validJSON)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	expectedResponse := &entity.PersonResponse{
		Gender:   "male",
		Fullname: "Mr John Doe",
		Address:  "123 Main St New York",
		Picture:  "https://example.com/profile.jpg",
	}

	if response.Gender != expectedResponse.Gender {
		t.Errorf("Expected gender: %s, but got: %s", expectedResponse.Gender, response.Gender)
	}
	if response.Fullname != expectedResponse.Fullname {
		t.Errorf("Expected fullname: %s, but got: %s", expectedResponse.Fullname, response.Fullname)
	}
	if response.Address != expectedResponse.Address {
		t.Errorf("Expected address: %s, but got: %s", expectedResponse.Address, response.Address)
	}
	if response.Picture != expectedResponse.Picture {
		t.Errorf("Expected picture: %s, but got: %s", expectedResponse.Picture, response.Picture)
	}
}
