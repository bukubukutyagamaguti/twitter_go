package usecase

import (
	"api/server/config"
	"api/server/domain"
	mock_usecase "api/server/usecase/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFollowInteractor_Update(t *testing.T) {
	type args struct {
		u domain.Follow
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockFollowRepository)
		args          args
		wantFollow    domain.Follow
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "正常値内にてフォローを増やす場合",
			prepareMockFn: func(mu *mock_usecase.MockFollowRepository) {
				mu.EXPECT().Update(domain.Follow{
					UserId:    1,
					FollowId:  2,
					CreatedAt: time.Now().Format(config.TimeFormat),
				}).Return(
					domain.Follow{
						Id:        1,
						UserId:    1,
						User:      &domain.User{},
						FollowId:  2,
						DeletedAt: gorm.DeletedAt{},
						CreatedAt: time.Now().Format(config.TimeFormat),
					}, nil)
			},
			args: args{
				u: domain.Follow{
					UserId:    1,
					FollowId:  2,
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
			},
			wantFollow: domain.Follow{
				Id:        1,
				UserId:    1,
				User:      &domain.User{},
				FollowId:  2,
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now().Format(config.TimeFormat),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockFollowRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewFollowInteractor(mu)
			gotFollow, err := uc.Update(tt.args.u)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantFollow, gotFollow)
		})
	}
}

// TODO deleted_atの値の入れ方を理解後対応
func TestFollowInteractor_DeleteById(t *testing.T) {
	type args struct {
		u domain.Follow
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockFollowRepository)
		args          args
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "正常値内にてフォローを削除する場合",
			prepareMockFn: func(mu *mock_usecase.MockFollowRepository) {
				mu.EXPECT().DeleteById(domain.Follow{
					Id:        1,
					UserId:    1,
					FollowId:  2,
					CreatedAt: time.Now().Format(config.TimeFormat),
				}).Return(nil)
			},
			args: args{
				u: domain.Follow{
					Id:        1,
					UserId:    1,
					FollowId:  2,
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
			mu := mock_usecase.NewMockFollowRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewFollowInteractor(mu)
			tt.assertion(t, uc.DeleteById(tt.args.u))
		})
	}
}

func TestFollowInteractor_SearchByFollowIdAndUserId(t *testing.T) {
	type args struct {
		query    string
		userId   int
		followId int
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockFollowRepository)
		args          args
		wantFollow    domain.Follow
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "user_id=1の場合のフォロワー一覧が取得できる場合",
			prepareMockFn: func(mu *mock_usecase.MockFollowRepository) {
				mu.EXPECT().WhereByUserIdAndFollowId("user_id = ? And follow_id = ?", 1, 2).Return(domain.Follow{
					Id:        1,
					UserId:    2,
					User:      &domain.User{},
					FollowId:  2,
					DeletedAt: gorm.DeletedAt{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				}, nil)
			},
			args: args{
				query:    "user_id = ? And follow_id = ?",
				userId:   1,
				followId: 2,
			},
			wantFollow: domain.Follow{
				Id:        1,
				UserId:    2,
				User:      &domain.User{},
				FollowId:  2,
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now().Format(config.TimeFormat),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_usecase.NewMockFollowRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewFollowInteractor(mu)
			gotFollow, err := uc.SearchByFollowIdAndUserId(tt.args.query, tt.args.userId, tt.args.followId)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantFollow, gotFollow)
		})
	}
}

func TestFollowInteractor_SearchFollowByUserId(t *testing.T) {
	type args struct {
		query string
		id    int
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *mock_usecase.MockFollowRepository)
		args          args
		wantFollows   domain.Follows
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "user_id=1の場合のフォロー一覧取得",
			prepareMockFn: func(mu *mock_usecase.MockFollowRepository) {
				mu.EXPECT().WhereById("user_id=?", 1).Return(domain.Follows{
					{
						Id:        1,
						UserId:    1,
						User:      &domain.User{},
						FollowId:  2,
						DeletedAt: gorm.DeletedAt{},
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
					{
						Id:        2,
						UserId:    1,
						User:      &domain.User{},
						FollowId:  3,
						DeletedAt: gorm.DeletedAt{},
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
					{
						Id:        3,
						UserId:    1,
						User:      &domain.User{},
						FollowId:  3,
						DeletedAt: gorm.DeletedAt{},
						CreatedAt: time.Now().Format(config.TimeFormat),
					},
				}, nil)
			},
			args: args{
				query: "user_id=?",
				id:    1,
			},
			wantFollows: domain.Follows{
				{
					Id:        1,
					UserId:    1,
					User:      &domain.User{},
					FollowId:  2,
					DeletedAt: gorm.DeletedAt{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
				{
					Id:        2,
					UserId:    1,
					User:      &domain.User{},
					FollowId:  3,
					DeletedAt: gorm.DeletedAt{},
					CreatedAt: time.Now().Format(config.TimeFormat),
				},
				{
					Id:        3,
					UserId:    1,
					User:      &domain.User{},
					FollowId:  3,
					DeletedAt: gorm.DeletedAt{},
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
			mu := mock_usecase.NewMockFollowRepository(ctrl)
			tt.prepareMockFn(mu)
			uc := NewFollowInteractor(mu)
			gotFollows, err := uc.SearchFollowByUserId(tt.args.query, tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantFollows, gotFollows)
		})
	}
}
