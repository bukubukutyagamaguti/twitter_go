package usecase

import (
	"api/server/config"
	"api/server/domain"
	mock_usecase "api/server/usecase/mock"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostInteractor_Add(t *testing.T) {
	type args struct {
		u domain.Post
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockPostRepository)
		args          args
		wantPost      domain.Post
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "正常値内にて投稿する場合",
			prepareMockFn: func(mu *mock_usecase.MockPostRepository) {
				mu.EXPECT().Store(domain.Post{
					Id:        1,
					UserId:    1,
					User:      &domain.User{},
					Message:   "テスト投稿だよ",
					DeletedAt: sql.NullTime{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				}).Return(domain.Post{
					Id:        1,
					UserId:    1,
					User:      &domain.User{},
					Message:   "テスト投稿だよ",
					DeletedAt: sql.NullTime{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				}, nil)
			},
			args: args{
				u: domain.Post{
					Id:        1,
					UserId:    1,
					User:      &domain.User{},
					Message:   "テスト投稿だよ",
					DeletedAt: sql.NullTime{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
			},
			wantPost: domain.Post{
				Id:        1,
				UserId:    1,
				User:      &domain.User{},
				Message:   "テスト投稿だよ",
				DeletedAt: sql.NullTime{},
				CreatedAt: time.Now().Format(config.TimeFormat),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockPostRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewPostInteractor(mu)
			gotPost, err := uc.Add(tt.args.u)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantPost, gotPost)
		})
	}
}

func TestPostInteractor_RelatedByUserId(t *testing.T) {
	type args struct {
		table string
		query string
		id    int
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockPostRepository)
		args          args
		wantPosts     domain.Posts
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "正常値内にてリクエストを出した場合",
			prepareMockFn: func(mu *mock_usecase.MockPostRepository) {
				mu.EXPECT().Related("User", "user_id = ?", 1).Return(domain.Posts{
					{
						Id:     1,
						UserId: 1,
						User: &domain.User{
							Id:        1,
							Name:      "テスト太郎",
							Email:     "test1@test.test",
							Password:  "password",
							Token:     sql.NullString{},
							Follows:   []*domain.Follow{nil},
							Posts:     []*domain.Post{nil},
							DeletedAt: sql.NullTime{},
							UpdatedAt: time.Now().Format(config.TimeFormat),
							CreatedAt: time.Now().Format(config.TimeFormat),
						},
						Message:   "テスト投稿だよ",
						DeletedAt: sql.NullTime{},
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
					{
						Id:     2,
						UserId: 2,
						User: &domain.User{
							Id:        2,
							Name:      "テスト次郎",
							Email:     "test2@test.test",
							Password:  "password",
							Token:     sql.NullString{},
							Follows:   []*domain.Follow{nil},
							Posts:     []*domain.Post{nil},
							DeletedAt: sql.NullTime{},
							UpdatedAt: time.Now().Format(config.TimeFormat),
							CreatedAt: time.Now().Format(config.TimeFormat),
						},
						Message:   "次郎の投稿だよ",
						DeletedAt: sql.NullTime{},
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
					{
						Id:     3,
						UserId: 2,
						User: &domain.User{
							Id:        2,
							Name:      "テスト次郎",
							Email:     "test2@test.test",
							Password:  "password",
							Token:     sql.NullString{},
							Follows:   []*domain.Follow{nil},
							Posts:     []*domain.Post{nil},
							DeletedAt: sql.NullTime{},
							UpdatedAt: time.Now().Format(config.TimeFormat),
							CreatedAt: time.Now().Format(config.TimeFormat),
						},
						Message:   "次郎のふたつめの投稿だよ",
						DeletedAt: sql.NullTime{},
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
				}, nil)
			},
			args: args{
				table: "User",
				query: "user_id = ?",
				id:    1,
			},
			wantPosts: domain.Posts{
				{
					Id:     1,
					UserId: 1,
					User: &domain.User{
						Id:        1,
						Name:      "テスト太郎",
						Email:     "test1@test.test",
						Password:  "password",
						Token:     sql.NullString{},
						Follows:   []*domain.Follow{nil},
						Posts:     []*domain.Post{nil},
						DeletedAt: sql.NullTime{},
						UpdatedAt: time.Now().Format(config.TimeFormat),
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
					Message:   "テスト投稿だよ",
					DeletedAt: sql.NullTime{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
				{
					Id:     2,
					UserId: 2,
					User: &domain.User{
						Id:        2,
						Name:      "テスト次郎",
						Email:     "test2@test.test",
						Password:  "password",
						Token:     sql.NullString{},
						Follows:   []*domain.Follow{nil},
						Posts:     []*domain.Post{nil},
						DeletedAt: sql.NullTime{},
						UpdatedAt: time.Now().Format(config.TimeFormat),
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
					Message:   "次郎の投稿だよ",
					DeletedAt: sql.NullTime{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
				{
					Id:     3,
					UserId: 2,
					User: &domain.User{
						Id:        2,
						Name:      "テスト次郎",
						Email:     "test2@test.test",
						Password:  "password",
						Token:     sql.NullString{},
						Follows:   []*domain.Follow{nil},
						Posts:     []*domain.Post{nil},
						DeletedAt: sql.NullTime{},
						UpdatedAt: time.Now().Format(config.TimeFormat),
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
					Message:   "次郎のふたつめの投稿だよ",
					DeletedAt: sql.NullTime{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockPostRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewPostInteractor(mu)
			gotPosts, err := uc.RelatedByUserId(tt.args.table, tt.args.query, tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantPosts, gotPosts)
		})
	}
}
