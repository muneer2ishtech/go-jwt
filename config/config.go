package config

// Configuration struct
type Config struct {
	Port       string
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

// LoadConfig loads configuration settings
func LoadConfig() *Config {
	return &Config{
		Port:       "8080",
		DBUsername: "gojwtuser",
		DBPassword: "gojwtpass",
		DBHost:     "127.0.0.1",
		DBPort:     "3306",
		DBName:     "gojwtdb",
	}
}
