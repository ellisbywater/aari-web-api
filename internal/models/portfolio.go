package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Portfolio struct {
	bun.BaseModel   `bun:"table:portfolios"`
	ID              string      `json:"-" bun:"type:uuid,pk,default:gen_random_uuid()"`
	CreatedAt       time.Time   `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time   `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	Name            string      `json:"name" bun:"name,notnull"`
	CapitalInvested float64     `json:"capital_invested" bun:"capital_invested,notnull"`
	Positions       []*Position `json:"positions" bun:"positions"`
}
