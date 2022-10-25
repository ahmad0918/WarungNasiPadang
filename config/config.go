package config

import (
	// "os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type DbConfig struct {
	Host 	 string
	Port 	 string
	Name 	 string
	User 	 string
	Password string
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

func (c Config) readConfigFile() Config {
	c.DbConfig = DbConfig{
		Host  : "localhost",
		Name  : "db_warung_nasi_padang",
		Port  : "5432",
		User  : "postgres",
		Password  : "123",
	}
	c.ApiConfig = ApiConfig{
		ApiPort: "8012",
		ApiHost: "localhost",
	}
	// c.DbConfig = DbConfig{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Name:     os.Getenv("DB_NAME"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// }
	// c.ApiConfig = ApiConfig{
	// 	ApiPort: os.Getenv("API_PORT"),
	// 	ApiHost: os.Getenv("API_HOST"),
	// }
	c.TokenConfig = TokenConfig{
		ApplicationName:     "ENIGMA",
		JwtSignatureKey:     "P@ssword",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: 60 * time.Second,
	}
	return c
}

func NewConfig() Config {
	cfg := Config{}
	return cfg.readConfigFile()
}