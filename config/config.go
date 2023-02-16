package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Debug    bool
}

var database *DatabaseConfig

func init() {
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("DB_USERNAME", "root")
	database.Password = gonv.GetStringEnv("PASSWORD", "")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("PORT", 3306)
	database.Database = gonv.GetStringEnv("DATABASE", "api-golang")
	database.Debug = gonv.GetBoolEnv("DEBUG", true)
}

func (this *DatabaseConfig) Url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", this.Username, this.Password, this.Host, this.Port, this.Database)
}

func GetDebug() bool {
	return database.Debug
}

func GetUrlDatabase() string {
	return database.Url()
}
