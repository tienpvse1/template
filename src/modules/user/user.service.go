package user

import (
	"context"
	"template/src/generated/sqlc"
)

type UserService struct {
	ctx     context.Context
	queries sqlc.Queries
}

func (service UserService) findById(id int32) (sqlc.User, error) {
	return service.queries.GetUser(service.ctx, id)
}
