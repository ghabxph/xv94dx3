package iface

import "sync"

type Endpoint interface {
	Create(wg *sync.WaitGroup)
}

type Database interface {
	Init()
}
