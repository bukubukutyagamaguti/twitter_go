package token

import (
	"api/server/domain"
	token_mock "api/server/interfaces/token/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTokenizerImpl_New(t *testing.T) {
	type args struct {
		user domain.User
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *token_mock.MockTokenHandler)
		args          args
		want          domain.Token
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "通常の手順にてToken発行",
			prepareMockFn: func(mu *token_mock.MockTokenHandler) {
				mu.EXPECT().Generate(1, "テスト太郎").Return(gomock.Any().String(), nil)
			},
			args: args{
				user: domain.User{
					Id:       1,
					Name:     "テスト太郎",
					Email:    "test@test.test",
					Password: "password",
				},
			},
			want:      domain.Token(gomock.Any().String()),
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := token_mock.NewMockTokenHandler(ctrl)
			tt.prepareMockFn(mu)
			tokenizer := NewTokenizerImpl(mu)
			got, err := tokenizer.New(tt.args.user)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTokenizerImpl_Verify(t *testing.T) {
	type args struct {
		token domain.Token
	}
	tests := []struct {
		name          string
		prepareMockFn func(mu *token_mock.MockTokenHandler)
		args          args
		assertion     assert.ErrorAssertionFunc
	}{
		{
			name: "tokenの内容に整合性が取れている場合",
			prepareMockFn: func(mu *token_mock.MockTokenHandler) {
				mu.EXPECT().VerityToken(gomock.Any().String()).Return(nil)
			},
			args:      args{token: domain.Token(gomock.Any().String())},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := token_mock.NewMockTokenHandler(ctrl)
			tt.prepareMockFn(mu)
			tokenizer := NewTokenizerImpl(mu)
			tt.assertion(t, tokenizer.Verify(tt.args.token))
		})
	}
}
