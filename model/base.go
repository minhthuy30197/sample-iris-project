package model

import (
	"context"
	"fmt"
	"sample-project/config"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// Kết nối database
func ConnectDB(user string, password string, database string, address string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     address,
	})

	return db
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil 
}

func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error{
	fmt.Println(q.FormattedQuery())
	return nil 
}

func SetupDatabase(db *pg.DB, config config.Config) {
	// Log query ra terminal
	db.AddQueryHook(dbLogger{})

	// Khởi tạo database
	err := MigrationDB(db, config)
	if err != nil {
		panic(err)
	}
}

func MigrationDB(db *pg.DB, config config.Config) error {
	// Tạo schema
	var schemas = []string{"blog"}
	for _, schema := range schemas {
		_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + ";")
		if err != nil {
			return err
		}
	}

	// Tạo bảng 
	/*---------------------- blog ----------------------*/
	var post Post 
	err := CreateTable(&post, db)
	if err != nil {
		return err 
	}

	return nil
}

func CreateTable(model interface{}, db *pg.DB) error {
	err := db.CreateTable(model, &orm.CreateTableOptions{
		Temp:          false,
		FKConstraints: true,
		IfNotExists:   true,
	})

	return err 
}

