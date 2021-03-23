package framework

import (
	"fmt"
	"github.com/ghabxph/xv94dx3/pkg/config"
)

type framework struct{}

func Init(c config.Config) {
	fmt.Println("Hello World!")

	r := router.Create(GofiberRouter{
		config: c.Endpoints["http"],
	})

	r.Get("/top/confirmed")
}
