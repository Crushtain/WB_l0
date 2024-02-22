package store

import (
	"WB_L0/config"
	"database/sql"
)

type Store struct {
	config *config.DB
	db     *sql.DB
}

func NewStore(config *config.DB) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	return nil
}

func (s *Store) Close() {

}
