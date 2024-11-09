package main

import (
	_ "github.com/proullon/ramsql/driver"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPass     string
	PostgresDatabase string
	PostgresInsecure bool
}

var (
	cfg = &Config{
		PostgresHost:     "localhost",
		PostgresPort:     "5432",
		PostgresUser:     "root",
		PostgresPass:     "root",
		PostgresDatabase: "cookify",
		PostgresInsecure: true,
	}
)
