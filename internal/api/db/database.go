package db

import (
	"cmd/api/internal/api/conf"
	"cmd/api/internal/domain/models/users"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	c := conf.New("conf/config.yaml")

	var d = DB

	d, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + c.DB.Host + " user=" + c.DB.User + " password=" + c.DB.Pwd + " dbname=" + c.DB.Database + " port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	DB = d
	d.AutoMigrate(&users.User{})
}

func AddUser(name string) {
	user := &users.User{Name: name}
	result := DB.Create(&user)
	fmt.Println(result)
}
