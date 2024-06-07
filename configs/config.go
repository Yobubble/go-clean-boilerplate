package configs

import "github.com/spf13/viper"

type Config struct {
	App App
	DB  Database
}
type App struct {
	Port int
}
type Database struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

// depends on config.yaml

func LoadConfig() Config {
	var config Config
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("./")

	configErr := viper.ReadInConfig()
	if configErr != nil {
		panic(configErr)
	}

	config.App.Port = viper.GetInt("app.server.port")
	config.DB.Host = viper.GetString("database.host")
	config.DB.Port = viper.GetInt("database.port")
	config.DB.Name = viper.GetString("database.name")
	config.DB.User = viper.GetString("database.user")
	config.DB.Password = viper.GetString("database.password")
	config.DB.SSLMode = viper.GetString("database.sslmode")

	return config
}
