package domain

type UserID string

func (u UserID) String() string {
	return string(u)
}

func NewUserID(id string) UserID {
	return UserID(id)
}
