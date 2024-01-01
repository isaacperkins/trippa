package db

import (
	"cmd/api/internal/conf"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	c := conf.GetConfig().DB

	d, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + c.Host + " user=" + c.User + " password=" + c.Pwd + " dbname=" + c.Database + " port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	DB = d
}

func Get() *gorm.DB {
	return DB
}
