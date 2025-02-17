package init

import (
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() *gorm.DB {
	if db == nil {
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?")
	}
	return db
}
