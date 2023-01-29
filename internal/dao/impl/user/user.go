package user

import (
	"context"
	"gin-admin-api/internal/model/sysUser"

	"gorm.io/gorm"
)

type UserDao struct {
	db           *gorm.DB
	sysUserQuery *sysUser.Query
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db:           db,
		sysUserQuery: sysUser.Use(db),
	}
}

func (u *UserDao) GetSysUserByUid(ctx context.Context, uid string) error {
	q := u.sysUserQuery.SysUser
	_, err := q.WithContext(ctx).Select(q.UID, q.Mobile).Where(q.UID.Eq(uid), q.IsDeleted.Is(false)).First()
	return err
}
