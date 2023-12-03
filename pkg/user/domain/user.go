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

	CreatedAt time.Time  `json:"created_at"                  bun:"created_at,type:datetime,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at"                  bun:"updated_at,type:datetime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"        bun:"deleted_at,type:datetime,soft_delete"`
}

func NewUser(name string) *User {
	return &User{
		ID:        NewUserID(uuid.New().String()),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
