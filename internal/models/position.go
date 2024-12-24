package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Position struct {
	bun.BaseModel   `bun:"table:positions,alias:p"`
	ID              string     `json:"-" bun:"type:uuid,pk,default:gen_random_uuid()"`
	CreatedAt       time.Time  `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time  `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	Ticker          string     `json:"ticker" bun:"ticker,notnull"`
	AssetType       string     `json:"asset_type" bun:"asset_type,notnull"`
	Bias            string     `json:"bias" bun:"bias,notnull"`
	Justification   string     `json:"justification" bun:"justification,notnull"`
	Expiration      time.Time  `json:"expiration" bun:"expiration"`
	CapitalInvested float64    `json:"capital_invested" bun:"capital_invested, notnull"`
	CapitalReturn   float64    `json:"capital_return" bun:"capital_return,notnull"`
	Portfolio       *Portfolio `bun:"rel:belongs-to,join:portfolio_id=id"`
	PortfolioID     string
}

type PositionCreate struct {
	Ticker          string    `json:"ticker"`
	AssetType       string    `json:"asset_type"`
	Bias            string    `json:"bias"`
	Justification   string    `json:"justification"`
	Expiration      time.Time `json:"expiration"`
	CapitalInvested float64   `json:"capital_invested"`
}

var _ bun.BeforeAppendModelHook = (*Position)(nil)

func (u *Position) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}

type PositionStore struct {
	db *bun.DB
}

func (ps *PositionStore) Create(ctx context.Context, position *PositionCreate) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := ps.db.NewInsert().Model(&position).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PositionStore) List(ctx context.Context, userID string) ([]*Position, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	positions := make([]*Position, 0)
	if err := ps.db.NewSelect().Model(&positions).Relation("User").Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}
	return positions, nil
}
