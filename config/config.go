package config


type Config struct {
	PostgresPort string
	PostgresPassword string
	PostgresDatabase string
	PostgresUser string
	PostgresHost string
}

func Load()  Config {

	var cfg Config
	cfg.PostgresPort ="5432"
	cfg.PostgresPassword ="0021"
	cfg.PostgresDatabase ="pulishla"
	cfg.PostgresUser ="nodirbek"
	cfg.PostgresHost ="localhost"

	return cfg
}