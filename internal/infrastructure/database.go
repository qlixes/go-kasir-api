package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPgsql(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=10",
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.DbName,
	)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	// Configure connection pool for Supabase
	db.SetMaxOpenConns(25)                 // Supabase free tier allows up to 50 connections
	db.SetMaxIdleConns(5)                  // Keep some idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Connection lifetime
	db.SetConnMaxIdleTime(2 * time.Minute) // Idle connection timeout

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	log.Println("Successfully connect database")

	return db, err
}
