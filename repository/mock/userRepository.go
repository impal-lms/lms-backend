package mock

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

type UserMockRepository struct {
	db map[int64]domain.User
}

func NewUserMockRepository() repository.UserRepository {
	adm, _ := domain.NewUser("admin", "admin@admin.com", "123456", domain.Admin)
	return &UserMockRepository{
		db: map[int64]domain.User{
			adm.Id: *adm,
		},
	}
}

func (repo UserMockRepository) GetUserByID(id int64) (*domain.User, error) {
	logrus.Println(repo.db, id)
	user, ok := repo.db[id]
	if ok {
		return &user, nil
	}

	return nil, errors.New("not found")
}

func (repo UserMockRepository) Authenticate(email, password string) (*domain.User, error) {
	for _, user := range repo.db {
		if user.Email == email {
			ok := user.CheckPasswordHash(password)
			if ok {
				return &user, nil
			}

			break
		}
	}

	return nil, errors.New("authentication failed")
}

func (repo *UserMockRepository) Save(user domain.User) error {
	_, exist := repo.db[user.Id]
	if exist {
		return errors.New("user with that id has already existed")
	}

	repo.db[user.Id] = user
	return nil
}
