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
}

var dbTest *DatabaseConfig

func init() {
	dbTest = &DatabaseConfig{}
	dbTest.Username = gonv.GetStringEnv("DB_USERNAME", "root")
	dbTest.Password = gonv.GetStringEnv("PASSWORD", "")
	dbTest.Host = gonv.GetStringEnv("HOST", "localhost")
	dbTest.Port = gonv.GetIntEnv("PORT", 3306)
	dbTest.Database = gonv.GetStringEnv("DATABASE", "api-golang")
}

func (this *DatabaseConfig) Url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", this.Username, this.Password, this.Host, this.Port, this.Database)
}

func GetUrlDatabase() string {
	return dbTest.Url()
}
