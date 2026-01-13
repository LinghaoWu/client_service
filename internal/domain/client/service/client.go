package service

import (
	"client_service/internal/domain/client/entity"
	"client_service/internal/domain/client/repository"
	"context"
)

type ClientDomain interface {
	CreateClient(c context.Context, e *entity.ClientDO) error
	GetClientByUSCC(c context.Context, uscc string) (*entity.ClientDO, error)
	GetClientByName(c context.Context, name string) (*entity.ClientDO, error)
}

type ClientDomainImpl struct {
	repo repository.ClientRepo
}

func NewClientDomain(repo repository.ClientRepo) ClientDomain {
	return &ClientDomainImpl{
		repo: repo,
	}
}

func (d *ClientDomainImpl) CreateClient(c context.Context, e *entity.ClientDO) error {
	return d.repo.CreateClient(c, e)
}

func (d *ClientDomainImpl) GetClientByUSCC(c context.Context, uscc string) (*entity.ClientDO, error) {
	return d.repo.GetClientByUSCC(c, uscc)
}

func (d *ClientDomainImpl) GetClientByName(c context.Context, name string) (*entity.ClientDO, error) {
	return d.repo.GetClientByName(c, name)
}
