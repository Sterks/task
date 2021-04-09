package config

import (
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	defaultHttpPort               = "8000"
	defaultHttpRTimeout           = 10 * time.Second
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultLimiterRPS             = 10
	defaultLimiterBurst           = 2
	defaultLimiterTTL             = 10 * time.Minute
	defaultVerificationCodeLength = 8
)

type Config struct {
	HTTP  HTTPConfig
	Mongo MongoConfig
}

type HTTPConfig struct {
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
}

type MongoConfig struct {
	URI      string
	User     string
	Password string
	Name     string `mapstructure:"databaseName"`
}

func Init(path string) (*Config, error) {
	defaultsParams()

	if err := parseConfigFie(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseEnv() error {
	if err := parseHostFromEnv(); err != nil {
		return err
	}
	return nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = viper.GetString("host")
	cfg.Mongo.User = viper.GetString("user")
	cfg.Mongo.Password = viper.GetString("password")
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	return nil
}

func parseConfigFie(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}

func defaultsParams() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultHttpMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
}

func parseHostFromEnv() error {
	viper.SetEnvPrefix("http")
	return viper.BindEnv("host")
}
