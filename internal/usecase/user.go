package usecase

import (
	"auth-service/internal/domain/dto"
	"auth-service/internal/domain/model"
	repository "auth-service/internal/repository/db"
	"auth-service/utils"
	"errors"
)

type UserUsecase interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Register(req dto.RegisterRequest) error {
	// Check if user already exists
	existing, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if existing != nil && existing.Email == req.Email {
		return errors.New("email already registered")
	}

	// Hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
	}

	err = u.userRepo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Login(req dto.LoginRequest) (string, error) {
	user, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return "", err
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
