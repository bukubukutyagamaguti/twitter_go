package usecase

import (
	"api/server/config"
	"api/server/domain"
	mock_token "api/server/interfaces/token/mock"
	mock_usecase "api/server/usecase/mock"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserInteractor_UserById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockUserRepository)
		args          args
		wantUser      *domain.User
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "ユーザーID=1が渡された場合はID=1のユーザーエンティティが返る",
			prepareMockFn: func(mu *mock_usecase.MockUserRepository) {
				mu.EXPECT().FindById(1).Return(domain.User{
					Id:        1,
					Name:      "テスト太郎",
					Email:     "test@test.test",
					Password:  "password",
					Token:     sql.NullString{},
					Follows:   []*domain.Follow{},
					Posts:     []*domain.Post{},
					DeletedAt: sql.NullTime{},
					UpdatedAt: time.Now().Format(config.TimeFormat),
					CreatedAt: time.Now().Format(config.TimeFormat),
				}, nil)
			},
			args: args{
				id: 1,
			},
			wantUser: &domain.User{
				Id:        1,
				Name:      "テスト太郎",
				Email:     "test@test.test",
				Password:  "password",
				Token:     sql.NullString{},
				Follows:   []*domain.Follow{},
				Posts:     []*domain.Post{},
				DeletedAt: sql.NullTime{},
				UpdatedAt: time.Now().Format(config.TimeFormat),
				CreatedAt: time.Now().Format(config.TimeFormat),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewUserInteractor(mu, mock_token.NewMockTokenizer(ctrl))
			gotUser, err := uc.UserById(tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantUser.Name, gotUser.Name)
		})
	}
}
func TestUserInteractor_Users(t *testing.T) {
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockUserRepository)
		wantUsers     domain.Users
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "ユーザ情報が正しく取得できた場合",
			prepareMockFn: func(mu *mock_usecase.MockUserRepository) {
				mu.EXPECT().FindAll().Return(domain.Users{
					{
						Id:   1,
						Name: "テスト太郎",
					},
					{
						Id:   2,
						Name: "テスト次郎",
					},
					{
						Id:   3,
						Name: "テスト次郎",
					},
				}, nil)
			},
			wantUsers: []domain.User{
				{
					Id:   1,
					Name: "テスト太郎",
				},
				{
					Id:   2,
					Name: "テスト次郎",
				},
				{
					Id:   3,
					Name: "テスト次郎",
				},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewUserInteractor(mu, mock_token.NewMockTokenizer(ctrl))
			gotUsers, err := uc.Users()
			tt.assertion(t, err)
			assert.Equal(t, tt.wantUsers, gotUsers)
		})
	}
}

func TestUserInteractor_Add(t *testing.T) {
	type args struct {
		u domain.User
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockUserRepository)
		args          args
		wantUser      domain.User
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "通常の手順通りにユーザを追加",
			prepareMockFn: func(mu *mock_usecase.MockUserRepository) {
				mu.EXPECT().Store(domain.User{
					Id:        1,
					Name:      "テスト太郎",
					Email:     "test@test.test",
					Password:  "password",
					Token:     sql.NullString{},
					Follows:   []*domain.Follow{},
					Posts:     []*domain.Post{},
					DeletedAt: sql.NullTime{},
					UpdatedAt: time.Now().Format(config.TimeFormat),
					CreatedAt: time.Now().Format(config.TimeFormat),
				}).Return(domain.User{
					Id:        1,
					Name:      "テスト太郎",
					Email:     "test@test.test",
					Password:  "password",
					Token:     sql.NullString{},
					Follows:   []*domain.Follow{},
					Posts:     []*domain.Post{},
					DeletedAt: sql.NullTime{},
					UpdatedAt: time.Now().Format(config.TimeFormat),
					CreatedAt: time.Now().Format(config.TimeFormat),
				}, nil)
			},
			args: args{
				u: domain.User{
					Id:        1,
					Name:      "テスト太郎",
					Email:     "test@test.test",
					Password:  "password",
					Token:     sql.NullString{},
					Follows:   []*domain.Follow{},
					Posts:     []*domain.Post{},
					DeletedAt: sql.NullTime{},
					UpdatedAt: time.Now().Format(config.TimeFormat),
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
			},
			wantUser: domain.User{
				Id:        1,
				Name:      "テスト太郎",
				Email:     "test@test.test",
				Password:  "password",
				Token:     sql.NullString{},
				Follows:   []*domain.Follow{},
				Posts:     []*domain.Post{},
				DeletedAt: sql.NullTime{},
				UpdatedAt: time.Now().Format(config.TimeFormat),
				CreatedAt: time.Now().Format(config.TimeFormat),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewUserInteractor(mu, mock_token.NewMockTokenizer(ctrl))
			gotUser, err := uc.Add(tt.args.u)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}

func TestUserInteractor_Update(t *testing.T) {
	type args struct {
		u domain.User
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockUserRepository)
		args          args
		wantUser      domain.User
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "通常通りにユーザをアップデート",
			prepareMockFn: func(mu *mock_usecase.MockUserRepository) {
				mu.EXPECT().Update(domain.User{
					Name:     "テスト太郎",
					Email:    "test@test.test",
					Password: "password",
				}).Return(domain.User{
					Id:        1,
					Name:      "テスト太郎",
					Email:     "test@test.test",
					Password:  "password",
					Token:     sql.NullString{},
					Follows:   []*domain.Follow{},
					Posts:     []*domain.Post{},
					DeletedAt: sql.NullTime{},
					UpdatedAt: time.Now().Format(config.TimeFormat),
					CreatedAt: time.Now().Format(config.TimeFormat),
				}, nil)
			},
			args: args{
				u: domain.User{
					Name:     "テスト太郎",
					Email:    "test@test.test",
					Password: "password",
				},
			},
			wantUser: domain.User{
				Id:        1,
				Name:      "テスト太郎",
				Email:     "test@test.test",
				Password:  "password",
				Token:     sql.NullString{},
				Follows:   []*domain.Follow{},
				Posts:     []*domain.Post{},
				DeletedAt: sql.NullTime{},
				UpdatedAt: time.Now().Format(config.TimeFormat),
				CreatedAt: time.Now().Format(config.TimeFormat),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewUserInteractor(mu, mock_token.NewMockTokenizer(ctrl))
			gotUser, err := uc.Update(tt.args.u)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}

func TestUserInteractor_DeleteById(t *testing.T) {
	type args struct {
		u domain.User
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockUserRepository)
		args          args
		wantUser      domain.User
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "通常通りユーザーを削除",
			prepareMockFn: func(mu *mock_usecase.MockUserRepository) {
				mu.EXPECT().DeleteById(domain.User{
					Id: 1,
				}).Return(nil)
			},
			args: args{
				u: domain.User{
					Id: 1,
				},
			},
			wantUser:  domain.User{},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewUserInteractor(mu, mock_token.NewMockTokenizer(ctrl))
			tt.assertion(t, uc.DeleteById(tt.args.u))
		})
	}
}

// func TestUserInteractor_Login(t *testing.T) {
// 	type fields struct {
// 		UserRepository UserRepository
// 		Tokenizer      token.Tokenizer
// 	}
// 	type args struct {
// 		login domain.LoginUser
// 	}
// 	tests := []struct {
// 		name      string
// 		fields    fields
// 		args      args
// 		want      domain.User
// 		want1     domain.Token
// 		assertion assert.ErrorAssertionFunc
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			interactor := &UserInteractor{
// 				UserRepository: tt.fields.UserRepository,
// 				Tokenizer:      tt.fields.Tokenizer,
// 			}
// 			got, got1, err := interactor.Login(tt.args.login)
// 			tt.assertion(t, err)
// 			assert.Equal(t, tt.want, got)
// 			assert.Equal(t, tt.want1, got1)
// 		})
// 	}
// }
