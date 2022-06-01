package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Persistor interface {
	Create(model interface{}) error
	FindBy(readModel interface{}, conditionals ...interface{}) error
	Search(readModels interface{}, conditionals ...interface{}) error
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(model interface{}) error {
	return r.db.Model(model).Create(model).Error
}

func (r *Repository) FindBy(readModel interface{}, conditionals ...interface{}) error {
	return r.db.First(readModel, conditionals...).Error
}

func (r *Repository) Search(readModels interface{}, conditionals ...interface{}) error {
	return r.db.Where(conditionals).Find(readModels).Error
}

var initDbConnection sync.Once
var repository *Repository

func NewRepository() *Repository {
	initDbConnection.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PWD")
		dbName := os.Getenv("DB_NAME")
		dns := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Etc/GMT",
			host,
			user,
			password,
			dbName,
			port,
		)

		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dns,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		if err != nil {
			log.Fatal(err.Error())
		}

		sqlDb, err := db.DB()
		if err != nil {
			log.Fatal(err.Error())
		}
		sqlDb.SetMaxIdleConns(5)
		sqlDb.SetMaxOpenConns(50)

		repository = &Repository{db: db}
	})

	return repository
}
