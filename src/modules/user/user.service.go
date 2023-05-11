package user

import (
	"template/src/common"
	"template/src/generated/sqlc"
)

type UserService common.Service

func (service UserService) findById(id int32) (sqlc.User, error) {
	return service.Imports.Queries.GetUser(service.Imports.Ctx, id)
}
