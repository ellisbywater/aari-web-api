package models

import "time"

type Order struct {
	ID             string     `json:"id" bun:"type:uuid,pk,default:gen_random_uuid()"`
	CreatedAt      time.Time  `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	OrderType      string     `json:"order_type"`
	Amount         float64    `json:"amount"`
	OrderDirection string     `json:"order_direction"`
	Portfolio      *Portfolio `json:"portfolio"`
	LimitPrice     float64    `json:"limit_price"`
	ProductType    string     `json:"product_type"`
	TimeInForce    string     `json:"tif"`
	Justification  string     `json:"justification"`
	Filled         bool       `json:"filled"`
}
