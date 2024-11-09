package config

type Config struct {
	FsRoot           string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPass     string
	PostgresDatabase string
	PostgresInsecure bool
	BunVerbose       bool
	ServerAddress    string
}

var (
	prod = &Config{
		PostgresHost:     "localhost",
		PostgresPort:     "5432",
		PostgresUser:     "root",
		PostgresPass:     "root",
		PostgresDatabase: "cookify",
		PostgresInsecure: true,
		BunVerbose:       true,
		ServerAddress:    "localhost:8080",
	}
	dev = &Config{
		PostgresHost:     "localhost",
		PostgresPort:     "5433",
		PostgresUser:     "root",
		PostgresPass:     "root",
		PostgresDatabase: "cookify",
		PostgresInsecure: true,
		BunVerbose:       true,
		ServerAddress:    "localhost:8080",
	}
	fs = &Config{
		FsRoot:        "data/fs",
		ServerAddress: "localhost:8080",
	}
)

func Default() *Config {
	return fs
}
