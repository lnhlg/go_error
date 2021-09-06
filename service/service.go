package service

import "go_error/model"

type Dao interface {
	GetUser(string) (*model.User, error)
}

type Service struct {
	db Dao
}

func NewService(db Dao) *Service {

	return &Service{db}

}
