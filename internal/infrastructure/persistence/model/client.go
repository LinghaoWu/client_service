package model

import (
	"client_service/internal/domain/client/entity"
	"time"
)

type ClientPO struct {
	ClientID  int64
	Name      string
	USCC      string
	Contact   string
	Phone     string
	Email     string
	Address   string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ClientDOToClientPO(do entity.ClientDO) *ClientPO {
	return &ClientPO{
		ClientID: do.ClientID,
		Name:     do.Name,
		USCC:     do.USCC,
		Contact:  do.Contact,
		Phone:    do.Phone,
		Email:    do.Email,
		Address:  do.Address,
		IsActive: true,
	}
}
