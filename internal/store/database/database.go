package database

import (
	"cmd/api/internal/conf"
	"cmd/api/internal/store/database/pgsql"
	"gorm.io/gorm"
)

func New(c *conf.Config) *gorm.DB {
	return pgsql.New(c)
}
