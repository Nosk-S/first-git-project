package config

import "github.com/joho/godotenv"

type Config struct {
	Database struct {
		User     string
		Password string
		Address  string
		Port     string
		Name     string
	}

	Server struct {
		Port string
	}
}

func New() Config {

	// env config
	envConfig, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}

	var cfg Config
	// db config
	cfg.Database.User = envConfig["DBUSER"]
	cfg.Database.Password = envConfig["DBPASS"]
	cfg.Database.Address = envConfig["DBADDRESS"]
	cfg.Database.Port = envConfig["DBPORT"]
	cfg.Database.Name = envConfig["DBNAME"]

	// Server config
	cfg.Server.Port = envConfig["SERVERPORT"]

	return cfg
}
