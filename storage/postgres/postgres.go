package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/storage"
)

type Storage struct {
	db  *sqlx.DB    // database connection
	log *log.Logger // logger

	userRepository storage.UserI // userRepository storage.UserI
	postRepository storage.PostI // postRepository storage.PostI
}

// CloseDB implements storage.StorageI
func (s *Storage) CloseDB() error {
	// Close the database connection.
	return s.db.Close()
}

// Post implements storage.StorageI
func (s *Storage) Post() storage.PostI {

	// if postRepository is not nil, return it
	if s.postRepository != nil {

		// return postRepository
		return s.postRepository
	}

	// if postRepository is nil, create a new one
	s.postRepository = newPostRepo(s.db, s.log)

	// return postRepository
	return s.postRepository
}

// User implements storage.StorageI
func (s *Storage) User() storage.UserI {

	// if userRepository is not nil, return it
	if s.userRepository != nil {

		// return userRepository
		return s.userRepository
	}

	// if userRepository is nil, create a new one
	s.userRepository = newUserRepo(s.db, s.log)

	// return userRepository
	return s.userRepository
}

// NewPostgres creates a new instance of the Postgres storage.
func NewPostgres(ctx context.Context, cfg *config.Config, log *log.Logger) (storage.StorageI, error) {

	// Create a connection string to connect to the database.
	postgresConnString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDB,
		cfg.PostgresPassword,
		cfg.PostgresSSLMode,
	)

	// log
	log.Println("Starting connection to the database...")

	// Connect to the database.
	db, err := sqlx.Connect("postgres", postgresConnString)

	// checking error
	if err != nil {
		log.Printf("Method: NewPostgres Comment: Connect to the database Error: %v", err)
		return nil, err
	}

	// Ping test
	err = db.Ping()

	// checking error
	if err != nil {
		log.Printf("Method: NewPostgres Comment: Ping test Error: %v", err)
		return nil, err
	}

	// log
	log.Println("Connection to the database is successful.")

	// Create a new instance of the Postgres storage and return it.
	return &Storage{
		db:  db,
		log: log,
	}, nil
}
