package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            string    `json:"-" bun:"type:uuid,pk,default:gen_random_uuid()"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	Email         string    `json:"email" bun:"email,notnull,unique"`
	Username      string    `json:"username" bun:"username,notnull,unique"`
	PasswordHash  string    `json:"password_hash" `
	IsActive      bool      `json:"is_active" bun:"is_active,default:'true'"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}

type UserStore struct {
	db *bun.DB
}

func (us *UserStore) Create(ctx context.Context, usercreate *UserCreate) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	// hash password

	hash, err := bcrypt.GenerateFromPassword([]byte(usercreate.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &User{
		Email:        usercreate.Email,
		Username:     usercreate.Username,
		PasswordHash: string(hash),
	}
	_, err = us.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
