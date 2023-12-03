package service

import (
	"context"

	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/pkg/user/cache"
	"github.com/veljkomatic/user-account/pkg/user/domain"
)

type CachedUserService interface {
	UserService
}

type cachedUserService struct {
	UserService
	userCache cache.UserCache
}

var _ CachedUserService = (*cachedUserService)(nil)

func NewCachedUserService(userService UserService, userCache cache.UserCache) CachedUserService {
	return &cachedUserService{
		UserService: userService,
		userCache:   userCache,
	}
}

func (u *cachedUserService) GetUser(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	user, err := u.userCache.Get(ctx, userID.String())
	if err != nil {
		log.Warn(ctx, "failed to get user from cache", log.Fields{
			"reason": err,
			"userID": userID,
		})
	}

	if user != nil {
		return user, nil
	}

	user, err = u.UserService.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err := u.userCache.Set(ctx, userID.String(), user); err != nil {
		return nil, err
	}

	return user, nil
}
