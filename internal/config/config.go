package config

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Name     string
}

func (dc *DatabaseConfig) DSN() string {
	return dc.User + ":" + dc.Password + "@tcp(" + dc.Host + ")/" + dc.Name
}

func LoadConfiguration() Config {
	return Config{
		Database: DatabaseConfig{
			User:     "root",
			Password: "root",
			Host:     "localhost",
			Name:     "happyTraveling",
		},
	}
}
