package databse

import (
	"context"
	"fmt"
	"sync"

	"github.com/gambitier/gocomm/config"
	"github.com/gambitier/gocomm/db/dal/authors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseRepo struct {
	pool           *pgxpool.Pool
	AuthorsQueries *authors.Queries
	// UserQueries    *authors.Queries
	// Add more queries here
}

var (
	dbRepo *DatabaseRepo
	once   sync.Once
)

// NewDatabaseRepo returns the singleton instance of Database
func NewDatabaseRepo(conf *config.Conf) (*DatabaseRepo, error) {
	var err error
	once.Do(func() {
		dbRepo, err = newDatabaseRepo(conf)
	})
	return dbRepo, err
}

func newDatabaseRepo(conf *config.Conf) (*DatabaseRepo, error) {
	pool, err := connectToDatabase(conf)
	if err != nil {
		return nil, err
	}

	return &DatabaseRepo{
		pool:           pool,
		AuthorsQueries: authors.New(pool),
	}, nil
}

func connectToDatabase(conf *config.Conf) (*pgxpool.Pool, error) {
	// Setup connection pooling
	config, err := pgxpool.ParseConfig(conf.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse DATABASE_URL: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}

func (db *DatabaseRepo) Close() {
	db.pool.Close()
}
