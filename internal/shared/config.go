package shared

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`

	Redis struct {
		Addr string        `mapstructure:"addr"`
		TTL  time.Duration `mapstructure:"ttl"`
	} `mapstructure:"redis"`

	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	// ENV override ─ APP_*
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = v.ReadInConfig()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
