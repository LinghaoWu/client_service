package clientsvc

import (
	"client_service/internal/domain/client/entity"
	"client_service/internal/domain/client/service"
	"client_service/pkg/snowflake"
	"context"
	"errors"
	"sync"
)

type Service interface {
	CreateClientSvc(c context.Context, e *entity.ClientDO) error
}

type ServiceImpl struct {
	client_domain service.ClientDomain
}

var (
	ServiceImplIns  *ServiceImpl
	ServiceImplOnce sync.Once
)

func GetServiceImpl(cd service.ClientDomain) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{
			client_domain: cd,
		}
	})
	return ServiceImplIns
}

func (s *ServiceImpl) CreateClientSvc(c context.Context, e *entity.ClientDO) error {
	if _, err := s.client_domain.GetClientByUSCC(c, e.USCC); err == nil {
		return errors.New("client with the given USCC already exists")
	}

	clientID, err := snowflake.GenerateClientID()
	if err != nil {
		return errors.New("failed to generate client ID: " + err.Error())
	}
	e.ClientID = clientID

	if err := s.client_domain.CreateClient(c, e); err != nil {
		return err
	}

	return nil
}
