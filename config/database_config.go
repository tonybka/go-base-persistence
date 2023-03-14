package config

// DatabaseConfig configurations of database connection
type DatabaseConfig struct {
	DBEndPoint string `mapstructure:"DB_ENDPOINT"`
	DBUserName string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}
