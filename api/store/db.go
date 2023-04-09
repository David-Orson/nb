package store

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"nautilus-billing.com/m/v2/store/model"
	"os"
)

func (s *Store) newDB() (err error) {
	host := os.Getenv("NB_DB_HOST")
	user := os.Getenv("NB_DB_USER")
	password := os.Getenv("NB_DB_PASSWORD")
	dbname := os.Getenv("NB_DB_NAME")
	port := os.Getenv("NB_DB_PORT")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London",
		host, user, password, dbname, port,
	)

	s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return
}

func (s *Store) migrations() (err error) {
	return s.DB.AutoMigrate(
		model.Account{},
	)
}
