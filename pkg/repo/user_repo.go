package repo

import (
	"errors"

	"github.com/Simonz2/Task_app/pkg/models"
	"github.com/Simonz2/Task_app/pkg/utils"
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
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// create a new user
func (r *UserRepo) CreateUser(username, password string) error {
	//check if user already exists
	_, err := r.GetUserByUsername(username)
	if err == nil {
		return errors.New("user already exists")
	}
	//validate password
	if !utils.ValidatePassword(password) {
		return errors.New("password does not meet the requirements")
	}
	user := &models.User{Username: username}
	if err := user.SetPassword(password); err != nil {
		return err
	}
	return r.db.Create(user).Error
}
