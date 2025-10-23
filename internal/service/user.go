package service

import "github.com/monforje/user-service/internal/repository"

type userService struct {
	repository repository.UserRepository
}

func newUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}
