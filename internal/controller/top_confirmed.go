package controller

import "github.com/ghabxph/xv94dx3/internal/entity/covidobservations"

type TopConfirmed struct{}

func (t *TopConfirmed) GetTopConfirmed(observation_date string, max_results uint) map[string]interface{} {

	covidStats := covidobservations.GetInstance()

	return covidStats.GetTopConfirmed(observation_date, max_results)
}
