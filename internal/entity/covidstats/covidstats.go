package covidstats

type covidStats struct{}

var instance *covidStats

func GetInstance() *covidStats {
	if instance == nil {
		instance = &covidStats{}
	}
	return instance
}

func (c *covidStats) GetTopConfirmed(observation_date string, max_results int) map[string]interface{} {
	sql := `
SELECT
	observation_date,
	country,
	SUM(confirmed) AS confirmed,
	SUM(deaths) AS deaths,
	SUM(recovered) AS recovered
FROM covid_statistics
WHERE observation_date = ?
ORDER BY confirmed DESC
GROUP BY country;
`
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
		"sql": sql,
	}
}