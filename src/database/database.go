package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DataAccessor struct {
}

func (DataAccessor) Connect() (db *sqlx.DB, err error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=UTC",
		"root",
		"saadmin@123",
		"localhost",
		3306,
		"shopping")

	db, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return db, nil
}

func (DataAccessor) Close(db *sqlx.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			fmt.Print(err.Error())
		}
	}
}
