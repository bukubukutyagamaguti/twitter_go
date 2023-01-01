package infrastructure

import (
	"api/server/config"
	"api/server/domain"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSqlMockHandler() (mysqlDb *gorm.DB, sqlMock sqlmock.Sqlmock, err error) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	mysqlDb, err = gorm.Open(
		mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}),
	)
	if err != nil {
		return nil, nil, err
	}

	return
}

func TestCreate(t *testing.T) {
	db, mock, err := NewSqlMockHandler()
	if err != nil {
		t.Fatal(err)
	}

	r := SQLHndoler{Conn: db}

	user := &domain.User{
		Name:      "テスト太郎くん",
		Email:     "test@test.com",
		Password:  "password",
		Token:     sql.NullString{},
		DeletedAt: sql.NullTime{},
		UpdatedAt: time.Now().Format(config.TimeFormat),
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
	// Mock設定
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `users` (`name`,`email`,`password`,`token`,`deleted_at`,`updated_at`,`created_at`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(user.Name, user.Email, user.Password, user.Token, user.DeletedAt, user.UpdatedAt, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(int64(user.Id), 1))
	mock.ExpectCommit()
	res := &domain.User{
		Name:      "テスト太郎くん",
		Email:     "test@test.com",
		Password:  "password",
		Token:     sql.NullString{},
		DeletedAt: sql.NullTime{},
		UpdatedAt: time.Now().Format(config.TimeFormat),
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
	// 実行
	if err := r.Create(&res).Error; err != nil {
		t.Fatal(err)
	}
}

func TestFind(t *testing.T) {
	db, mock, err := NewSqlMockHandler()
	if err != nil {
		t.Fatal(err)
	}

	r := SQLHndoler{Conn: db}

	user := &domain.User{
		Id:        1,
		Name:      "テスト太郎くん",
		Email:     "test@test.com",
		Password:  "password",
		UpdatedAt: time.Now().Format(config.TimeFormat),
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
	id := 1
	// Mock設定
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE `users`.`id` = ?")).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(user.Id, user.Name))

	res := &domain.User{}
	// 実行
	if err := r.Find(&res, id).Error; err != nil {
		t.Fatal(err)
	}
	if res.Id != user.Id || res.Name != user.Name {
		t.Errorf("不一致 %v", res)
	}
}
func TestDelete(t *testing.T) {
	db, mock, err := NewSqlMockHandler()
	if err != nil {
		t.Fatal(err)
	}

	r := SQLHndoler{Conn: db}

	user := &domain.User{
		Id:       1,
		Name:     "テスト太郎くん",
		Email:    "test@test.com",
		Password: "password",
		Token:    sql.NullString{},
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now().Format(config.TimeFormat),
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
	// Mock設定
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `users` WHERE `users`.`id` = ?")).
		WithArgs(user.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res := &domain.User{
		Id:       1,
		Name:     "テスト太郎くん",
		Email:    "test@test.com",
		Password: "password",
		Token:    sql.NullString{},
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now().Format(config.TimeFormat),
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
	// 実行
	if err := r.Delete(&res).Error; err != nil {
		t.Fatal(err)
	}

}
func TestWhere(t *testing.T) {
	db, mock, err := NewSqlMockHandler()
	if err != nil {
		t.Fatal(err)
	}

	r := SQLHndoler{Conn: db}

	user := &domain.User{
		Id:   1,
		Name: "テスト太郎くん",
	}
	id := 1
	// Mock設定
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE `id`=?")).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(user.Id, user.Name))

	res := &domain.User{}
	// 実行
	if err := r.Where(&res, "`id`=?", id).Error; err != nil {
		t.Fatal(err)
	}
	if res.Id != user.Id || res.Name != user.Name {
		t.Errorf("不一致 %v", res)
	}
}

func TestPreloadAndWhere(t *testing.T) {
	db, mock, err := NewSqlMockHandler()
	if err != nil {
		t.Fatal(err)
	}

	r := SQLHndoler{Conn: db}

	post := &domain.Post{
		Id:     2,
		UserId: 1,
		User: &domain.User{
			Id:        1,
			Name:      "テスト太郎くん",
			Email:     "test@test.com",
			Password:  "password",
			UpdatedAt: time.Now().Format(config.TimeFormat),
			CreatedAt: time.Now().Format(config.TimeFormat),
		},
		Message:   "",
		DeletedAt: sql.NullTime{},
		CreatedAt: "",
	}
	id := 1
	// Mock設定
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `posts` WHERE user_id = ?")).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id))

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE users.id = ?")).
		WithArgs(post.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(post.Id))

	res := &domain.Post{}
	// 実行
	if err := r.PreloadAndWhere(&res, "User", "user_id = ?", id).Error; err != nil {
		t.Fatal(err)
	}
}
