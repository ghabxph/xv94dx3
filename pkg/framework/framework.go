package framework

import (
	"sync"
	"github.com/ghabxph/xv94dx3/pkg/config"
)

func Init(config config.Config) {

	var wg sync.WaitGroup

	// Run endpoints
	for _, endpoint := range config.Endpoints {
		wg.Add(1)
		go endpoint.Create(&wg)
	}

	// Initialize database
	config.Database.Init()

	// Wait for all endpoint threads to end)
	wg.Wait()
}
