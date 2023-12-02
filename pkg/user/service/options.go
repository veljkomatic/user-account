package service

import "github.com/veljkomatic/user-account/pkg/user/domain"

type UserOption func(user *domain.User)

// Add here options that are used to create a new user, options are not mandatory like CreateUserParams
// example:
// func SharedUser(shared bool) UserOption {
//	return func(user *domain.User) {
//		user.Shared = shared
//	}
//}
