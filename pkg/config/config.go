package config

import (
	"github.com/ghabxph/xv94dx3/pkg/config/iface"
)

type Config struct {
	Endpoints map[string]iface.Endpoint
	Database iface.Database
}
