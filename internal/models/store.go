package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var QueryTimeoutDuration = time.Second * 5

type Store struct {
	Users interface {
		Create(ctx context.Context, usercreate *UserCreate) error
	}
	Positions interface {
		Create(context.Context, *Position) error
	}
}

func InitStore(db *bun.DB) *Store {
	return &Store{
		Users:     &UserStore{db},
		Positions: &PositionStore{db},
	}
}
