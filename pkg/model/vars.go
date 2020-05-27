package model

import (
	"time"

	"github.com/go-pg/pg"
)

var (
	loggingMode int
	pgdb        *pg.DB
)

type Model struct {
	Id int64
}
type Blameable struct {
	UsersId int64 `pg:"on_delete:RESTRICT, notnull"`
}

type Deleatables struct {
	CreatedAt time.Time `pg:",notnull, default:now()"`
	// UpdatedAt time.Time
	// DeletedAt time.Time `pg:"soft_delete"`
}
