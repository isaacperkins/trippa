package store

import (
	users "cmd/api/internal/domain/users/model"
	database "cmd/api/internal/store/pgsql"
	"fmt"
)

func UsersAdd(u *users.User) *users.User {
	db := database.Get()
	r := db.Create(&u)
	if r.Error != nil {
		fmt.Println(r.Error)
	}
	return u
}

func UsersGet() *[]users.User {
	db := database.Get()
	u := []users.User{}
	db.Find(&u)
	return &u
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

/*
func Migrate() {
	db := database.Get()
	db.AutoMigrate(&users.User{})
}
*/
