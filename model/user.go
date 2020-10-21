package model

import (
	"fmt"

	"github.com/init-new-world/Go_API_learning/pkg/constvar"

	"github.com/go-playground/validator/v10"

	"github.com/init-new-world/Go_API_learning/pkg/auth"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;not null;unique" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *User) NewRecord() bool {
	if d := DB.Self.Where("username = ?", u.Username).First(&u).Error; d != nil {
		return true
	}
	return false
}

func (u *User) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *User) Delete() error {
	return DB.Self.Where("username = ?", u.Username).Delete(&User{}).Error
}

func (u *User) Update() error {
	return DB.Self.Save(u).Error
}

func (u *User) Get() error {
	d := DB.Self.Where("username = ?", u.Username).First(&u)
	return d.Error
}

func ListUser(username string, offset, limit int) ([]*User, uint, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*User, 0)
	var count uint

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
