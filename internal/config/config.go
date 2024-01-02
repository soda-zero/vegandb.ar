package config

type Config struct {
	DatabaseConnectionString string
}

const (
	Dev  = "local"
	Prod = "production"
)

var ConnectionStrings = map[string]string{
	Dev:  "file:db.sqlite",
	Prod: "your_production_connection_string",
}

func GetConfig(environment string) Config {
	return Config{
		DatabaseConnectionString: ConnectionStrings[environment],
	}
}

