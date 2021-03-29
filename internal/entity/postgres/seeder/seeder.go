package seeder

import (
    "log"
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
    "github.com/ghabxph/xv94dx3/internal/entity/postgres"
    "github.com/araddon/dateparse"
    "io"
    "io/ioutil"
    "encoding/csv"
    "strings"
    "strconv"
    "os"
)

type PostgresSeeder struct {}

func (p *PostgresSeeder) RunSeeder() {

    log.Println("Running seeder.")

    db := pg.Connect(postgres.GetInstance().Options())
    defer db.Close()

    // 1. Creates the schema
    createSchema(db)

    // 2. Read csv
    readCsvAndPopulateTable(db)
}

func createSchema(db *pg.DB) {

    log.Println("Creating DB Schema")

    err := db.Model(&postgres.CovidObservations{}).CreateTable(&orm.CreateTableOptions{})
    if err != nil {
        panic("PostgreSQL schema creation failed: " + err.Error())
    }
}

func readCsvAndPopulateTable(db *pg.DB) {
    log.Println("Reading CSV")

    fname := os.Args[1][12:]
    //in, err := ioutil.ReadFile("covid_19_data.csv")
    in, err := ioutil.ReadFile(fname)
    if err != nil {
        panic(err)
    }

    r := csv.NewReader(strings.NewReader(string(in)))
    r.Read()

    for item, err := r.Read(); err != io.EOF; item, err = r.Read() {
        sno, _:= strconv.ParseUint(item[0], 10, 64)
        observationDate, e1 := dateparse.ParseAny(item[1])
        if e1 != nil {
            panic(e1)
        }
        lastUpdate, _ := dateparse.ParseAny(item[4])
        confirmed, _ := strconv.ParseFloat(item[5], 64)
        deaths, _ := strconv.ParseFloat(item[6], 64)
        recovered, _ := strconv.ParseFloat(item[7], 64)
        covidStats := &postgres.CovidObservations{
            SNo: sno,
            ObservationDate: observationDate,
            ProvinceState: item[2],
            CountryRegion: item[3],
            LastUpdate: lastUpdate,
            Confirmed: uint64(confirmed),
            Deaths: uint64(deaths),
            Recovered: uint64(recovered),
        }
        db.Model(covidStats).Insert()
        log.Println(covidStats)
    }

    log.Println("read through csv done")
}
