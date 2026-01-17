package client_repo

import (
	"client_service/internal/domain/client/entity"
	"client_service/internal/domain/client/repository"
	"client_service/internal/infrastructure/ent"
	"client_service/internal/infrastructure/ent/clientschema"
	"client_service/internal/infrastructure/persistence/model"
	"context"
)

type ClientRepoImpl struct {
	db *ent.Client
}

func NewClientRepo(db *ent.Client) repository.ClientRepo {
	return &ClientRepoImpl{
		db: db,
	}
}

func (r *ClientRepoImpl) CreateClient(c context.Context, e *entity.ClientDO) error {
	po := model.ClientDOToClientPO(*e)

	_, err := r.db.ClientSchema.Create().
		SetClientID(po.ClientID).
		SetName(po.Name).
		SetUSCC(po.USCC).
		SetContact(po.Contact).
		SetPhone(po.Phone).
		SetEmail(po.Email).
		SetAddress(po.Address).
		SetIsActive(po.IsActive).
		Save(c)

	return err
}

func (r *ClientRepoImpl) GetClientByUSCC(c context.Context, uscc string) (*entity.ClientDO, error) {
	client, err := r.db.ClientSchema.Query().
		Where(clientschema.IsActive(true)).
		Where(clientschema.USCC(uscc)).
		Only(c)

	if err != nil {
		return nil, err
	}

	return &entity.ClientDO{
		ClientID:  client.ClientID,
		Name:      client.Name,
		USCC:      client.USCC,
		Contact:   client.Contact,
		Phone:     client.Phone,
		Email:     client.Email,
		Address:   client.Address,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}

func (r *ClientRepoImpl) GetClientByName(c context.Context, name string) (*entity.ClientDO, error) {
	client, err := r.db.ClientSchema.Query().
		Where(clientschema.IsActive(true)).
		Where(clientschema.Name(name)).
		Only(c)

	if err != nil {
		return nil, err
	}

	return &entity.ClientDO{
		ClientID:  client.ClientID,
		Name:      client.Name,
		USCC:      client.USCC,
		Contact:   client.Contact,
		Phone:     client.Phone,
		Email:     client.Email,
		Address:   client.Address,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
