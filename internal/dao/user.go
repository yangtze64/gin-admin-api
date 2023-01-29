package dao

import "context"

var userDao IUserDao

type IUserDao interface {
	GetSysUserByUid(ctx context.Context, uid string) error
}

func RegisterUserDao(i IUserDao) {
	userDao = i
}
