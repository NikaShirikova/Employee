package postgresql

import (
	"Employee/internal/module"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=empl sslmode=disable"
	dbase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "empl.",
			SingularTable: false,
		},
	})
	if err != nil {
		panic("failed to connect to database")
	}

	dbase.Migrator().AutoMigrate(&module.Employee{}, &module.Passport{})

	return dbase
}

func GetDB() *gorm.DB {
	if db == nil {
		db = Init()
		var sleep = time.Duration(1)
		for db == nil {
			sleep = sleep * 2
			fmt.Println("Database is unavailable.")
			time.Sleep(sleep * time.Second)
			db = Init()
		}
	}
	return db
}
