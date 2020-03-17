package core

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"go-playground/core/settings"
)

var DB *pg.DB

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// Create db connection
func DbConnect() *pg.DB {
	s := settings.GetSettings()
	DB = pg.Connect(&pg.Options{
		Addr:     s.GetString("database.address"),
		User:     s.GetString("database.user"),
		Password: s.GetString("database.password"),
		Database: s.GetString("database.name"),
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

// Enable printing queries which this library generates
func EnableDbQueryLogger() {
	db := GetDb()
	db.AddQueryHook(dbLogger{})
}
