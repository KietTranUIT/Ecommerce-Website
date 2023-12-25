package repository

import (
	"fmt"
	"user-service/internal/core/port/repository"
	"user-service/internal/infra/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	*gorm.DB
}

func NewDB() (repo repository.Database, err error) {
	var conf config.ConfigDB
	conf, err = config.LoadConfigDB("/home/kiettran/IT/ecommerce-project/internal/infra/config/")

	if err != nil {
		return
	}

	var db *gorm.DB
	db, err = connectDB(conf)
	if err != nil {
		return
	}

	return database{db}, err
}

func connectDB(conf config.ConfigDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBPassword,
		conf.DBAddress,
		conf.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db database) GetDB() *gorm.DB {
	return db.DB
}

func (db database) Close() error {
	dbInstance, err := db.DB.DB()

	if err != nil {
		return err
	}

	return dbInstance.Close()
}
