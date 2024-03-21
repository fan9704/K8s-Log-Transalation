package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	SqlSession *gorm.DB
)
func Initialize(dbConfig string) (*gorm.DB, error) {
	var err error
	SqlSession, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	return SqlSession, err
}