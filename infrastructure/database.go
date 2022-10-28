package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLHndoler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SQLHndoler {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	containerName := os.Getenv("DB_CONTAINER_NAME")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, containerName, port, dbName)
	conn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	sqlHandler := new(SQLHndoler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SQLHndoler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}
func (handler *SQLHndoler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}
func (handler *SQLHndoler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}
func (handler *SQLHndoler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}
func (handler *SQLHndoler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}
func (handler *SQLHndoler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}
func (handler *SQLHndoler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}
func (handler *SQLHndoler) Where(out interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...).Find(out)
}
func (handler *SQLHndoler) Joins(query string, args ...interface{}) *gorm.DB {
	return handler.Conn.Joins(query, args)
}
func (handler *SQLHndoler) Distinct(query ...interface{}) *gorm.DB {
	return handler.Conn.Distinct(query)
}
