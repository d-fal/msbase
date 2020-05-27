package model

import (
	"fmt"
	"log"
	"msbase/pkg/libs"
	"reflect"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/logrusorgru/aurora"
	_ "github.com/paulmach/orb"
)

func DbModel() *pg.DB {

	return DbConnect()
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
	fmt.Println(aurora.Red("----------------"))

}

func CreateModels() {
	pgdb := DbConnect()

	for _, model_ := range []interface{}{
		// &User{}, &Country{}, # all the models should be added
	} {

		err := pgdb.CreateTable(model_, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
		})
		if err != nil {
			log.Printf("Error with %s : %s", reflect.TypeOf(model_).Elem().Name(), err)

		} else {
			log.Printf("Model %s created!", reflect.TypeOf(model_).Elem().Name())
		}
	}

}

func DbConnect() *pg.DB {
	if pgdb != nil {
		return pgdb
	}
	// err, _ := pgdb.Exec("SELECT 1")
	// if err != nil {
	pgdb = pg.Connect(&pg.Options{
		User:     "tmuser",
		Password: "tmuser",
		Database: "tm",
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 5432),
	})
	switch GetLoggingMode() {
	case libs.LoggingModeVerbose:
		pgdb.AddQueryHook(dbLogger{})

	}
	// }

	return pgdb
}
