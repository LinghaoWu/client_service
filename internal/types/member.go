package types

import "client_service/internal/domain/member/entity"

type MemberReq struct {
	MemberID int64  `json:"member_id"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type MemberResp struct {
	MemberID int64  `json:"member_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func MemberReqToMemberDO(req *MemberReq) *entity.MemberDO {
	return &entity.MemberDO{
		MemberID: req.MemberID,
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
		Role:     req.Role,
	}
}

func MemberDOToMemberResp(do *entity.MemberDO) *MemberResp {
	return &MemberResp{
		MemberID: do.MemberID,
		Name:     do.Name,
		Phone:    do.Phone,
		Email:    do.Email,
		Role:     do.Role,
	}
}
