package repository

import (
	"context"
	"github.com/veljkomatic/user-account/common/ptr"
	"github.com/veljkomatic/user-account/common/storage/sql/postgres"
	"github.com/veljkomatic/user-account/pkg/user/domain"
)

type ReadOnlyUserRepository interface {
	GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
}

type UserRepository interface {
	ReadOnlyUserRepository

	SaveUser(ctx context.Context, user *domain.User) error
}

var _ UserRepository = (*repository)(nil)

type repository struct {
	idb postgres.IDB
}

func NewUserRepository(idb postgres.IDB) UserRepository {
	return &repository{
		idb: idb,
	}
}

func (r *repository) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	var user domain.User
	if err := r.idb.NewSelect().Model(&user).
		Where("user.id = ?", userID).
		Scan(ctx); err != nil {
		return nil, postgres.HandleAndWrapErr(err, "failed to get user by id")
	}

	return ptr.From(user), nil
}

func (r *repository) SaveUser(ctx context.Context, user *domain.User) error {
	if _, err := r.idb.NewInsert().Model(user).Exec(ctx); err != nil {
		return postgres.HandleAndWrapErr(err, "failed to save user")
	}

	return nil
}
