package main

import (
	c "github.com/ghabxph/xv94dx3/pkg/config"
	"github.com/ghabxph/xv94dx3/pkg/framework"
	"github.com/ghabxph/xv94dx3/pkg/config/iface"
	"github.com/ghabxph/xv94dx3/internal/entity"
	"github.com/ghabxph/xv94dx3/internal/endpoint"
	"os"
	"strconv"
)

func main() {

	// Init http port and db port
	http_port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	db_port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	// Setup the framework and run
	framework.Init(c.Config{
		Endpoints: map[string]iface.Endpoint {
			"http": &endpoint.Http{Port: http_port},
		},
		Database: &entity.PostgreSQLDatabase{
			Port: db_port,
			DbName: os.Getenv("DB_NAME"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	})
}
