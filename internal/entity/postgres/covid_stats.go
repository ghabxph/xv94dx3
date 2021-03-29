package postgres

import (
	//"log"
	"time"
	"github.com/go-pg/pg/v10"
)

// SNo,ObservationDate,Province/State,Country/Region,Last Update,Confirmed,Deaths,Recovered
type CovidStats struct {
	SNo uint64 `pg:",notnull"`
	ObservationDate time.Time
	ProvinceState string
	CountryRegion string
	LastUpdate time.Time
	Confirmed uint64 `pg:",notnull,use_zero"`
	Deaths uint64 `pg:",notnull,use_zero"`
	Recovered uint64 `pg:",notnull,use_zero"`
}

func (c *CovidStats) GetTopConfirmed(observation_date string, max_results uint) map[string]interface{} {

	db := pg.Connect(GetInstance().Options())
	defer db.Close()

	var covidStats []CovidStats

	db.Model(&covidStats).Column("country_region").
		ColumnExpr("SUM(confirmed) AS confirmed").
		ColumnExpr("SUM(deaths) AS deaths").
		ColumnExpr("SUM(recovered) AS recovered").
		Where("? = ?", pg.Ident("observation_date"), observation_date).
		Group("country_region").
		Order("confirmed DESC").
		Limit(int(max_results)).
		Select()

	countries := make([]map[string]interface{}, 0)

	for _, country := range covidStats {
		countries = append(countries, map[string]interface{}{
			"country": country.CountryRegion,
			"confirmed": country.Confirmed,
			"deaths": country.Deaths,
			"recovered": country.Recovered,
		})
	}

	return map[string]interface{}{
		"observation_date": observation_date,
		"countries": countries,
	}
}
