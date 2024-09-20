package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

var db *pgx.Conn

func Init() error {
	connUrl, exist := os.LookupEnv("POSTGRES_CONN")
	if !exist {
		return errors.New("POSTGRES_CONN doesn't exist!")
	}
	fmt.Println(connUrl)
	var err error
	db, err = pgx.Connect(context.Background(), connUrl)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to ping database: %v\n", err)
	}

	db.Exec(
		context.Background(),
		`CREATE TYPE tender_status AS ENUM (
    	'Created',
    	'Published',
    	'Closed');`)

	db.Exec(
		context.Background(),
		`CREATE TYPE tender_service_type AS ENUM (
		'Construction',
		'Delivery',
		'Manufacture');`)

	db.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS tenders (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		title VARCHAR(100) NOT NULL,
		description VARCHAR(500),
		service_type tender_service_type NOT NULL,
		status tender_status DEFAULT 'Created',
		organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
		version INT DEFAULT 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`)

	db.Exec(context.Background(),
	`CREATE TABLE IF NOT EXISTS user_tenders(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID REFERENCES employee(id) ON DELETE CASCADE,
	tender_id UUID REFERENCES tenders(id) ON DELETE CASCADE);`)

	return nil
}
