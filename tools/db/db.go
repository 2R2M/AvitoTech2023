package db

import (
	"avitoTech/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DefaultDatabasePort         = 1433
	DefaultDatabaseMaxOpenConns = 5
)

func NewDB(c *config.DB) (*sqlx.DB, error) {
	conf := setDefault(c)

	connString := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Port,
		conf.Name,
	)
	fmt.Println(connString)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()

		return nil, err
	}

	log.Println("Database sqlx connection established")

	return db, nil
}

func setDefault(c *config.DB) (conf config.DB) {

	if c != nil {
		conf = *c
	}

	if conf.Port == 0 {
		conf.Port = DefaultDatabasePort
	}

	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = DefaultDatabaseMaxOpenConns
	}

	return
}
