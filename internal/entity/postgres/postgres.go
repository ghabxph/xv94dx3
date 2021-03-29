package postgres

import (
	"fmt"
    "github.com/go-pg/pg/v10"
	"github.com/ghabxph/xv94dx3/pkg/seeder"
	"os"
)

type Database struct {
	Host string
	Port int
	DbName string
	DbUsername string
	DbPassword string
	Seeder seeder.Seeder
}

var instance *Database

func (d *Database) Init() {

	if (instance == nil) {
		instance = d
	}

	if len(os.Args) != 2 {
		return // No parameter given, thus closing...
	}

	if os.Args[1][0:12] != "--runSeeder=" {
		return // Exit if parameter is not --runSeedder
	}

	d.Seeder.RunSeeder()
}

func (d *Database) Options() *pg.Options {
	opt, err := pg.ParseURL(fmt.Sprintf(
        "postgres://%s:%s@%s:%d/%s?sslmode=disable",
        d.DbUsername,
        d.DbPassword,
        d.Host,
        d.Port,
        d.DbName,
    ))
    if err != nil {
        panic(err)
    }
	return opt
}

func GetInstance() *Database {
	return instance
}
