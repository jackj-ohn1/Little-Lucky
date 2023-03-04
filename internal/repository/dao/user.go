package dao

import (
	"errors"
	"gorm.io/gorm"
	"temp/internal/repository"
)

func Insert(account, email string) error {
	var user = repository.User{
		Account: account,
		Email:   email,
	}
	
	return repository.DB.Table("users").Create(&user).Error
}

// 没有找到返回true
func Check(account string) (error, bool) {
	var user repository.User
	if err := repository.DB.Table("users").Where("account=?", account).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, true
		}
		return err, false
	}
	return nil, false
}
