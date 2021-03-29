package covidstats

type CovidStats interface {
	GetTopConfirmed(observation_date string, max_results uint) map[string]interface{}
}

var instance CovidStats

func CreateOnce(_instance CovidStats) {
	if instance != nil {
		return
	}
	instance = _instance
}

func GetInstance() CovidStats {
	return instance
}
