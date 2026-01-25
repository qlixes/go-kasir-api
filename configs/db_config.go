package configs

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func LoadDB() *gorm.DB {
	dsn := GetEnv("DB_CONNECTION_URL")
	dbOnce.Do(func() {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = db
	})

	return dbInstance
}
