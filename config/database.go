package config

import (
	"fmt"
	"quanlykhachsan/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "monorail.proxy.rlwy.net"
	port     = 34913
	user     = "postgres"
	password = "Da332Fb5ea4ACAf2BD3E1ef-*c16gc-5"
	dbName   = "railway"
)

func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
