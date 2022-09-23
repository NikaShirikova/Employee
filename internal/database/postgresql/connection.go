package postgresql

import (
	"Employee/internal/module"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Host        string
	Username    string
	Password    string
	DBName      string
	SSLMode     string
	TablePrefix string
}

var db *gorm.DB

func Init(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	dbase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix,
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}

	errorMigrate := dbase.Migrator().AutoMigrate(&module.Employee{}, &module.Passport{}, &module.Department{}, &module.Company{})
	if errorMigrate != nil {
		return dbase, errorMigrate
	}
	return dbase, nil
}

//func GetDB(cfg Config) *gorm.DB {
//	if db == nil {
//		db, _ = Init(cfg)
//		var sleep = time.Duration(1)
//		for db == nil {
//			sleep = sleep * 2
//			fmt.Println("Database is unavailable.")
//			time.Sleep(sleep * time.Second)
//			db, _ = Init(cfg)
//		}
//	}
//	return db
//}
