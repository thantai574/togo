package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port       string     `json:"port" mapstructure:"port"`
	Env        string     `json:"env" mapstructure:"env"`
	JwtKey     string     `json:"jwt_key" mapstructure:"jwt_key"`
	DBPostgres DBPostgres `json:"postgres" mapstructure:"postgres"`
}

type DBPostgres struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	UserName string `json:"user_name" mapstructure:"user_name"`
	Port     string `json:"port"  mapstructure:"port"`
	Password string `json:"password" mapstructure:"password"`
	DB       string `json:"db" mapstructure:"db"`
	SSLMode  string `json:"ssl_mode" mapstructure:"ssl_mode"`
	TimeZone string `json:"time_zone"  mapstructure:"time_zone"`
}

// LoadTestConfig load config for running tests
func LoadConfig(configPath, name string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	viper.SetConfigName(name)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	result := &Config{}
	err = viper.Unmarshal(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
