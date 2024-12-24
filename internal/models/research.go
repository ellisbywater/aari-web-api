package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Research struct {
	bun.BaseModel   `bun:"table:research"`
	ID              string    `json:"-" bun:"type:uuid,pk,default:gen_random_uuid()"`
	CreatedAt       time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	Topic           string    `json:"topic" bun:"topic,notnull"`
	Products        []string  `json:"products" bun:"products,notnull"`
	PointsOfThought []string  `json:"points_of_thought" bun:"points_of_thought"`
}
