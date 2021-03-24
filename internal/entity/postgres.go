package entity

type PostgreSQLDatabase struct {
	Port int
	DbName string
	DbUsername string
	DbPassword string
}

func (p *PostgreSQLDatabase) Init() {

}
