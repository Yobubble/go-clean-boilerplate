package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App App
	DB  Database
	Client Client
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

type Client struct {
	AllowedClient []string
}

func LoadConfig() Config {
	var config Config
	godotenv.Load()

	viper.SetConfigName(os.Getenv("NODE_ENV"))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

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
	config.Client.AllowedClient = viper.GetStringSlice("client.allowed")
	
	return config
}
