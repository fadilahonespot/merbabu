package utils

import "gorm.io/gorm"

func Paginate(page, length int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case length > 50:
			length = 50
		case length <= 0:
			length = 10
		}

		offset := (page - 1) * length
		return db.Offset(offset).Limit(length)
	}
}
