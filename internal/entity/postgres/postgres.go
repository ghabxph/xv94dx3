package postgres

import (
	"fmt"
    "github.com/go-pg/pg/v10"
	"github.com/ghabxph/xv94dx3/pkg/seeder"
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

	// TODO: Logic for database seeder. Seeder must only run when certain parameter is set
	// --runSeeder
	// --runSeeder=covid_19_data.csv
	// Must read --runSeeder
	// Second iteration: must read --runSeeder and covid_19_data.csv
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
