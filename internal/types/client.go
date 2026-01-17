package types

import "client_service/internal/domain/client/entity"

type ClientReq struct {
	ClientID int64  `json:"client_id"`
	Name     string `json:"name" binding:"required"`
	USCC     string `json:"uscc" binding:"required"`
	Contact  string `json:"contact"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type ClientResp struct {
	ClientID int64  `json:"client_id"`
	Name     string `json:"name"`
	USCC     string `json:"uscc"`
	Contact  string `json:"contact"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

func ClientReqToClientDO(req *ClientReq) *entity.ClientDO {
	return &entity.ClientDO{
		ClientID: req.ClientID,
		Name:     req.Name,
		USCC:     req.USCC,
		Contact:  req.Contact,
		Phone:    req.Phone,
		Email:    req.Email,
		Address:  req.Address,
	}
}

func ClientDOToClientResp(do *entity.ClientDO) *ClientResp {
	return &ClientResp{
		ClientID: do.ClientID,
		Name:     do.Name,
		USCC:     do.USCC,
		Contact:  do.Contact,
		Phone:    do.Phone,
		Email:    do.Email,
		Address:  do.Address,
	}
}
