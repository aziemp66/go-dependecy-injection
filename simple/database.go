package simple

type Database struct {
	Name string
}

type DatabasePostgresql Database
type DatabaseMongodb Database

func NewDatabasePostgresql() *DatabasePostgresql {
	return &DatabasePostgresql{Name: "Postgresql"}
}

func NewDatabaseMongodb() *DatabaseMongodb {
	return &DatabaseMongodb{Name: "Mongodb"}
}

type DatabaseRepository struct {
	DatabasePostgresql *DatabasePostgresql
	DatabaseMongodb    *DatabaseMongodb
}

func NewDatabaseRepository(databasePostgresql *DatabasePostgresql, databaseMongodb *DatabaseMongodb) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgresql: databasePostgresql,
		DatabaseMongodb:    databaseMongodb,
	}
}
