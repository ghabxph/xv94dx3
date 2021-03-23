package main

import (
	c "github.com/ghabxph/xv94dx3/pkg/config"
	"github.com/ghabxph/xv94dx3/internal/framework"
	"os"
	"strconv"
)

func main() {

	// Init http port and db port
	http_port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	db_port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	// Setup the framework and run
	framework.Init(c.Config{
		Endpoints: map[string]c.Endpoint {
			"http": c.HttpEndpoint{Port: http_port},
		},
		Databases: map[string]c.Database {
			"postgres": c.PostgreSQLDatabase{
				Port: db_port,
				DbName: os.Getenv("DB_NAME"),
				DbUsername: os.Getenv("DB_USERNAME"),
				DbPassword: os.Getenv("DB_PASSWORD"),
			},
		},
	})
}
