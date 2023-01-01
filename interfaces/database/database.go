//go:generate mockgen -source=database.go -package=mock_database -destination=./mock/database.go
package database

import "gorm.io/gorm"

type SqlHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, interface{}, ...interface{}) *gorm.DB
	PreloadAndWhere(interface{}, string, string, ...interface{}) *gorm.DB
}
