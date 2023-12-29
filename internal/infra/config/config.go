package config

import (
	"github.com/spf13/viper"
)

type ConfigDB struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBAddress  string `mapstructure:"DB_ADDRESS"`
	DBName     string `mapstructure:"DB_NAME"`
}

// receive parameter is path to file.env and return ConfigDB, error
func LoadConfigDB(path string) (conf ConfigDB, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("database")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)

	conf.DBAddress = "web-database.mysql.database.azure.com"
	conf.DBUser = "kiettran"
	conf.DBPassword = "nhom20@123456"
	conf.DBName = "ecommerce"

	err = nil
	return
}
