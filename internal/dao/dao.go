package dao

import (
	"gin-admin-api/internal/dao/impl/user"

	"gorm.io/gorm"
)

func User(db *gorm.DB) IUserDao {
	if userDao == nil {
		RegisterUserDao(user.NewUserDao(db))
	}
	return userDao
}
