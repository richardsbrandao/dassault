package db

import "gopkg.in/pg.v4"

func Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "dassault",
		SSL:      false,
		PoolSize: 13000,
	})
}
