package interfaceservice

import "Web_Service/domain/entity"

type RandomUserService interface {
	GetRandomUser() (*entity.PersonResponse, error)
}
