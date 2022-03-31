package manager

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	SqlDb() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func (i *infra) SqlDb() *sqlx.DB {
	return i.db
}

func NewInfra(dataSourceName string) Infra {
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &infra{
		db: conn,
	}
}
