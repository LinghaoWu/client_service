package repository

import (
	"client_service/internal/domain/client/entity"
	"context"
)

type ClientRepo interface {
	CreateClient(c context.Context, e *entity.ClientDO) error
	GetClientByUSCC(c context.Context, uscc string) (*entity.ClientDO, error)
	GetClientByName(c context.Context, name string) (*entity.ClientDO, error)
}
