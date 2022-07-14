package logic

import "gin-admin-api/app/logic/impl/member"

func init() {
	RegisterMember(member.New())
}
