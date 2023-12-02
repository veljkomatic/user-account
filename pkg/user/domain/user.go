package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:user"`

	ID   UserID `json:"id" bun:"type:uuid,notnull,pk"`
	Name string `json:"name" bun:"type:varchar(255),notnull"`

	CreatedAt time.Time  `json:"created_at" bun:"type:timestamptz,notnull,default:now()"`
	UpdatedAt time.Time  `json:"updated_at" bun:"type:timestamptz,notnull,default:now()"`
	DeleteAt  *time.Time `json:"delete_at" bun:"type:timestamptz"`
}

func NewUser(name string) *User {
	return &User{
		ID:        NewUserID(uuid.New().String()),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
