package postgres

import (
	"WB_l0/config"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const (
	defaultMaxPoolSize  int32 = 1
	defaultConnAttempts       = 10
	defaultConnTimeout        = time.Second
)

type Postgres struct {
	Pool         *pgxpool.Pool
	Builder      squirrel.StatementBuilderType
	maxPoolSize  int32 //max count of connections, int32 for type identity in pgxpool
	connAttempts int   //count of connection attempts to the database
	connTimeout  time.Duration
}

func New(url string) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}
	// Custom options
	//for _, opt := range opts {
	//	opt(pg)
	//}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) //allow to place params to db request by $
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = pg.maxPoolSize // max count of connections for poolConfig as in Postgres{}

	//trying to connect
	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			break
		}

		//нужен logger

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}
	if err != nil {
		return nil, err
	}

	return pg, nil

}

func GetConnString(db *config.DB) string {

	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password='%s' sslmode=disable search_path=%s",
		db.Host,
		db.Port,
		db.User,
		db.Password,
		db.Schema,
	)

	return connString
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
