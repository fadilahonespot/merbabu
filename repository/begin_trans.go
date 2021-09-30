package repository

import "gorm.io/gorm"

func (e *defaultRepository) BeginTrans() *gorm.DB {
	return e.db.Begin()
}
