package core

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

var DB *pg.DB

// Create db connection
func DbConnect() *pg.DB {
	DB = pg.Connect(&pg.Options{
		User:     "gpuser",
		Password: "qppassword",
		Database: "gpdatabase",
	})
	return DB
}

// Create db tables by models if they doesn't exist yet
func CreateSchema(db *pg.DB, models []interface{}) error {
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Return current db connection
func GetDb() *pg.DB {
	return DB
}

// Close current db connection
func CloseDatabase() error {
	err := DB.Close()
	if err != nil {
		return err
	}
	return nil
}
