package databases

import "gorm.io/gorm"

type Database interface {
	GetDB() *gorm.DB
}
