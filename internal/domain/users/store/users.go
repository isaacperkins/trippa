package store

import (
	users "cmd/api/internal/domain/users/model"
	database "cmd/api/internal/store/pgsql"
	"fmt"
)

func AddUser(u *users.User) *users.User {
	db := database.Get()
	result := db.Create(&u)
	fmt.Println(result)
	return u
}

func Migrate() {
	db := database.Get()
	db.AutoMigrate(&users.User{})
}
