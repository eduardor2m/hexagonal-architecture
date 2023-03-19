package usecases

import "github.com/eduardor2m/hexagonal-architecture/src/internal/domain/models"

type UserRepository interface {
	Save(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type UserUseCase struct {
	userRepo UserRepository
}

func NewUserUseCase(userRepo UserRepository) *UserUseCase {
	return &UserUseCase{userRepo}
}

func (u *UserUseCase) CreateUser(name, email string) error {
	user := &models.User{Name: name, Email: email}

	err := u.userRepo.Save(user)

	if err != nil {
		return err
	}

	return nil

}

func (u *UserUseCase) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.userRepo.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
