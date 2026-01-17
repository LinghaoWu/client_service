package entity

type MemberDO struct {
	MemberID int64
	Name     string
	Phone    string
	Email    string
	Role     string
}

func (do *MemberDO) IsValidMember() bool {
	if len(do.Email) == 0 || len(do.Phone) == 0 {
		return false
	}

	return true
}
