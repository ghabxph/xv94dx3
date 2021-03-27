package postgres

import "time"

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

func (c *CovidStats) GetTopConfirmed(observation_date string, max_results int) map[string]interface{} {
	return map[string]interface{}{
		"observation_date": "yyyy-mm-dd",
		"countries": []map[string]interface{}{
			{
				"country":   "Mainland China",
				"confirmed": 15,
				"deaths":    2,
				"recovered": 7,
			},
			{
				"country":   "Italy",
				"confirmed": 10,
				"deaths":    3,
				"recovered": 5,
			},
		},
	}
}
