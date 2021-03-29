package covidobservations

type CovidObservations interface {
	GetTopConfirmed(observation_date string, max_results uint) map[string]interface{}
}

var instance CovidObservations

func CreateOnce(_instance CovidObservations) {
	if instance != nil {
		return
	}
	instance = _instance
}

func GetInstance() CovidObservations {
	return instance
}
