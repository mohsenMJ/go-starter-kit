package repositories

import (
	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/models"
	"github.com/mohsenMj/go-starter-kit/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user models.User) models.User
	Update(user models.User) models.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
}

type userConnection struct {
	// DB *gorm.DB
}

func newUserRepository() UserRepository {
	return &userConnection{}
}

func (r *userConnection) Insert(user models.User) models.User {
	user.Password, _ = utils.HashPassword(user.Password)
	database.DB.Save(&user)
	return user
}

func (r *userConnection) Update(user models.User) models.User {
	user.Password, _ = utils.HashPassword(user.Password)
	database.DB.Save(&user)
	return user
}
func (r *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := database.DB.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
func (r *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return database.DB.Where("email = ?").Take(&user)
}
func (r *userConnection) FindByEmail(email string) models.User {
	var user models.User
	database.DB.Where("email = ?").Take(&user)
	return user

}
func (r *userConnection) ProfileUser(userID string) models.User {
	var user models.User
	database.DB.Find(&user, userID)
	return user
}
