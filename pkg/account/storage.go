package account

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
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
	CreatedAt    time.Time
	ModifiedAt   *time.Time
}

const (
	userTable        = "`user`"
	userAccountTable = "`user_account`"
)

var (
	userAccountCols = []string{"id", "user_sub", "open_id", "union_id", "nick", "source", "avatar", "email", "access_token", "scope", "refresh_token", "expired_at", "raw", "created_at", "modified_at"}
)

type RdsStorage struct {
	db *sql.DB
}

func (r *RdsStorage) SaveUser(ctx context.Context, user *User) error {
	onError := func(err error) error {
		return errors.Wrap(err, "RdsStorage.SaveUser")
	}

	builder := sq.Insert(userTable).Columns("sub", "primary_account").Values(user.Sub, user.PrimaryAccount)
	query, args, err := builder.ToSql()
	if err != nil {
		return onError(err)
	}

	if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
		return onError(err)
	}

	return nil
}

func (r *RdsStorage) SaveUserAccount(ctx context.Context, account *UserAccount) error {
	onError := func(err error) error {
		return errors.Wrap(err, "RdsStorage.SaveUserAccount")
	}

	builder := sq.Replace(userAccountTable).Columns(userAccountCols...).
		Values(nil, account.UserSub, account.OpenId, account.UnionId, account.Nick, account.Source, account.Avatar, account.Email, account.AccessToken, account.Scope, account.RefreshToken, account.ExpiredAt, account.Raw, nil, nil)
	query, args, err := builder.ToSql()
	if err != nil {
		return onError(err)
	}

	if result, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return onError(err)
	} else {
		if id, err := result.LastInsertId(); err != nil {
			return onError(err)
		} else {
			account.Id = id
		}
	}

	return nil
}

func (r *RdsStorage) GetUserAccount(ctx context.Context, source, openid string) (*UserAccount, error) {
	onError := func(err error) (*UserAccount, error) {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "RdsStorage.GetUserAccount")
	}

	builder := sq.Select(userAccountCols...).From(userAccountTable).Where("source=? AND open_id=?", source, openid)
	query, args, err := builder.ToSql()
	if err != nil {
		return onError(err)
	}

	var account UserAccount
	if err = r.db.QueryRowContext(ctx, query, args...).
		Scan(&account.Id, &account.UserSub, &account.OpenId, &account.UnionId, &account.Nick, &account.Source, &account.Avatar,
			&account.Email, &account.AccessToken, &account.Scope, &account.RefreshToken, &account.ExpiredAt, &account.Raw, &account.CreatedAt, &account.ModifiedAt);
		err != nil {
		return onError(err)
	}

	return &account, nil
}

func newRdsStorage(db *sql.DB) Storage {
	return &RdsStorage{db: db}
}

type Storage interface {
	SaveUser(ctx context.Context, user *User) error
	SaveUserAccount(ctx context.Context, account *UserAccount) error
	GetUserAccount(ctx context.Context, source, openid string) (*UserAccount, error)
}

type mockStorage struct {
	userAccount map[string]map[string]*UserAccount
	user        map[string]*User
}

func (s *mockStorage) SaveUser(ctx context.Context, user *User) error {
	s.user[user.Sub] = user
	return nil
}

func (s *mockStorage) SaveUserAccount(ctx context.Context, account *UserAccount) error {
	if _, ok := s.userAccount[account.Source]; !ok {
		s.userAccount[account.Source] = map[string]*UserAccount{}
	}
	s.userAccount[account.Source][account.OpenId] = account
	return nil
}

func (s *mockStorage) GetUserAccount(ctx context.Context, source, openid string) (*UserAccount, error) {
	return s.userAccount[source][openid], nil
}

func newMockStorage() Storage {
	return &mockStorage{
		userAccount: make(map[string]map[string]*UserAccount),
		user:        make(map[string]*User),
	}
}
