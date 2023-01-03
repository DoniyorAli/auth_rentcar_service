package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	homeDB *sqlx.DB
}

func InitDB(psqlUrl string) (*Postgres, error) {
	var err error

	psqlDB, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		return nil, err
	} 

	return &Postgres{
		homeDB: psqlDB,
	}, nil
}