package store

import (
	users "cmd/api/internal/domain/users/model"
	database "cmd/api/internal/store/pgsql"
	"fmt"
	"gorm.io/gorm"
)

func UsersGet(d *gorm.DB) *[]users.User {
	u := []users.User{}
	d.Find(&u)
	return &u
}

func UserGet(d *gorm.DB, id string) *users.User {
	u := users.User{}
	d.First(&u, id)
	return &u
}

func UsersAdd(u *users.User) *users.User {
	db := database.Get()
	r := db.Create(&u)
	if r.Error != nil {
		fmt.Println(r.Error)
	}
	return u
}

func UsersEdit(u *users.User) *users.User {
	db := database.Get()
	r := db.Save(&u)
	if r.Error != nil {
		fmt.Println(r.Error)
	}
	return u
}

func UsersDelete(id string) string {
	db := database.Get()
	r := db.Delete(&users.User{}, id)
	if r.Error != nil {
		fmt.Println(r.Error)
	}
	return "success"
}

func Migrate(d *gorm.DB) string {
	d.AutoMigrate(&users.User{})
	return "success"
}
