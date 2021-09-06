package dao

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type SqlServerDB struct {
	*gorm.DB
}

func NewSqlServerDB(dsn string) (*SqlServerDB, error) {

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err

	}

	return &SqlServerDB{db}, nil
}
