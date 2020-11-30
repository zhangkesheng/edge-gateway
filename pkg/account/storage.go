package account

import (
	"context"
	"time"
)

type User struct {
	// 用户唯一标识
	Sub string
	// 主账号ID
	PrimaryAccount int64
	CreateAt       string
}

type UserAccount struct {
	Id           int64
	UserSub      string
	OpenId       string
	UnionId      string
	Nick         string
	Source       string
	Avatar       string
	Email        string
	AccessToken  string
	Scope        string
	RefreshToken string
	ExpiredAt    int64
	Raw          interface{}
	CreateAt     time.Time
	ModifiedAt   time.Time
}

type Storage interface {
	SaveUser(ctx context.Context, user *User) error
	SaveUserAccount(ctx context.Context, account *UserAccount) error
	GetUserAccount(ctx context.Context, source, openid string) (*UserAccount, error)
}
