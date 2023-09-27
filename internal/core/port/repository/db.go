package repository

import (
	"io"

	"gorm.io/gorm"
)

type Database interface {
	io.Closer
	GetDB() *gorm.DB
}
