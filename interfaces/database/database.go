package database

import "gorm.io/gorm"

type SqlHandler interface {
	Exec(string, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Raw(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, interface{}, ...interface{}) *gorm.DB
	Joins(string, ...interface{}) *gorm.DB
	Distinct(...interface{}) *gorm.DB
	PreloadAndWhere(interface{}, string, string, ...interface{}) *gorm.DB
}
