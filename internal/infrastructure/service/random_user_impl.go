package service

import (
	"Web_Service/domain/entity"
	interfaceservice "Web_Service/domain/interface_service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type randomUserService struct{}

func NewRandomUserService() interfaceservice.RandomUserService {
	return &randomUserService{}
}

func (s *randomUserService) GetRandomUser() (*entity.PersonResponse, error) {
	body, err := GetRandomUser()

	if err != nil {
		return nil, err
	}

	response, err := ParseJSON(body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetRandomUser() ([]byte, error) {
	resp, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func ParseJSON(body []byte) (*entity.PersonResponse, error) {
	var result entity.RandomUserResponse
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Results) == 0 {
		return nil, fmt.Errorf("NO RESULTS FOUND FOR RANDOM USER")
	}

	user := result.Results[0]
	response := &entity.PersonResponse{
		Gender:   user.Gender,
		Fullname: fmt.Sprintf("%s %s %s", user.Name.Title, user.Name.First, user.Name.Last),
		Address:  fmt.Sprintf("%s %s", user.Location.Street.Name, user.Location.City),
		Picture:  user.Picture.Large,
	}

	return response, nil

}
