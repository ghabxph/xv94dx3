package config

type Config struct{
	Endpoints map[string]Endpoint
	Databases map[string]Database
}

type Endpoint interface{}
type Database interface{}

type HttpEndpoint struct{
	Port int
}

type PostgreSQLDatabase struct{
	Port int
	DbName string
	DbUsername string
	DbPassword string
}
