package mygrations

import "database/sql"

type Deps struct {
	MasterDB    *sql.DB
	SlaveDB     *sql.DB
	SomeFactory interface{} //example
}
