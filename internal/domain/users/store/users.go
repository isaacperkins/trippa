package store

import (
	users "cmd/api/internal/domain/users/model"
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

func UsersAdd(d *gorm.DB, u *users.User) *users.User {
	r := d.Create(&u)
	if r.Error != nil {
		fmt.Println(r.Error)
	}
	return u
}

func UsersEdit(d *gorm.DB, u *users.User) *users.User {
	r := d.Save(&u)
	if r.Error != nil {
		fmt.Println(r.Error)
	}
	return u
}

func UsersDelete(d *gorm.DB, id string) string {
	r := d.Delete(&users.User{}, id)
	if r.Error != nil {
		fmt.Println(r.Error)
	}
	return "success"
}

func Migrate(d *gorm.DB) string {
	d.AutoMigrate(&users.User{})
	return "success"
}
