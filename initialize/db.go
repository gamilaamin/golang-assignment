package initialize

import (
	"fmt"
	"golang-assignment/entity"
	"golang-assignment/global"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         225,     // string The length of the type field by default
		SkipInitializeWithVersion: false,   // According to the version automatic configuration
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.Pet{},
	)
	if err != nil {
		panic(fmt.Errorf("register table failed: %s \n", err))
	}
	log.Printf("register table success\n")
}
