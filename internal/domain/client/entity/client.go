package entity

import "time"

type ClientDO struct {
	ClientID  int64
	Name      string
	USCC      string
	Contact   string
	Phone     string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (do *ClientDO) IsValidClient() bool {
	if len(do.USCC) == 0 || len(do.USCC) > 18 {
		return false
	}
	return true
}
