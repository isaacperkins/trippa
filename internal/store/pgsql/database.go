package pgsql

import (
	"cmd/api/internal/conf"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New(c *conf.Config) *gorm.DB {

	d, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + c.DB.Host + " user=" + c.DB.User + " password=" + c.DB.Pwd + " dbname=" + c.DB.Database + " port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	DB = d
	return d
}

func Get() *gorm.DB {
	return DB
}
