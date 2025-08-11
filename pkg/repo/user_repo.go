package repo

import (
	"errors"

	"github.com/Simonz2/Task_app/pkg/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// check if user exists by username and return it
func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username=?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
func (r *UserRepo) CreateUser(username, password string) error {
	err := r.db.Where("username=?", username).First(&models.User{}).Error
	if err == nil {
		return errors.New("user already exists")
	}
	user := &models.User{Username: username}
	if err := user.SetPassword(password); err != nil {
		return err
	}
	return r.db.Create(user).Error
}
