package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AppConfig struct {
	Server *ServerConfig
	Jwt    *JwtConfig
}

type ServerConfig struct {
	Address string
}

type JwtConfig struct {
	SecretKey string
}

func NewApplicationConfig() *AppConfig {
	config := &AppConfig{}

	viper.SetConfigName("application.yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Read config file error: %s", err.Error())
		panic(err)
	}

	config.Server = &ServerConfig{
		Address: ":" + viper.GetString("server.address"),
	}
	config.Jwt = &JwtConfig{
		SecretKey: viper.GetString("jwt.secret-key"),
	}

	return config
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../../data.db"), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Connect to DB error: %s", err.Error())
		panic(err)
	}

	return db, nil
}