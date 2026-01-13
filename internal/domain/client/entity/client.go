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

func (c *ClientDO) IsValidClient() bool {
	if len(c.USCC) == 0 || len(c.USCC) > 18 {
		return false
	}
	return true
}
