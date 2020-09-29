package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string   `mapstructure:"env"`
	StackTrace  bool     `mapstructure:"stack_trace"`
	Server      Server   `mapstructure:"server"`
	LogLevel    string   `mapstructure:"log_level"`
	LogFormat   string   `mapstructure:"log_format"`
	Database    Database `mapstructure:"database"`
	EncryptoKey string   `mapstructure:"encrypto_key"`
}

type Database struct {
	Driver       string `mapstructure:"driver"`
	InstanceName string `mapstructure:"instance_name"`
	Host         string `mapstructure:"host"`
	Port         uint   `mapstructure:"port"`
	SSLMode      bool   `mapstructure:"ssl_mode"`
	DBName       string `mapstructure:"dbname"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	MaxLifetime  string `mapstructure:"max_lifetime"`
	MaxIdleConn  int    `mapstructure:"max_idle_conn"`
	MaxOpenConn  int    `mapstructure:"max_open_conn"`
	LogMode      bool   `mapstructure:"log_mode"`
	Timeout      string `mapstructure:"connect_timeout"`
}

type Server struct {
	GrpcPort string `mapstructure:"grpc_port"`
}

func NewConfig(configPath string) (config Config, err error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	/* default */
	v.SetDefault("env", "local")
	v.SetDefault("stack_trace", false)
	v.SetDefault("log_level", "INFO")
	v.SetDefault("log_format", "console")
	v.SetDefault("http.port", "8080")
	v.SetDefault("database.driver", "postgres")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.ssl_mode", false)
	v.SetDefault("database.instance_name", "")
	v.SetDefault("database.dbname", "kbc")
	v.SetDefault("database.user", "root")
	v.SetDefault("database.password", "silkrode1234")
	v.SetDefault("database.max_lifetime", "1h")
	v.SetDefault("database.connect_timeout", "10s")
	v.SetDefault("database.max_idle_conn", 2)
	v.SetDefault("database.max_open_conn", 5)
	v.SetDefault("database.log_mode", false)
	v.SetDefault("auth.identity_address", "127.0.0.1:17486")
	v.SetDefault("auth.token_ttl", int64(43200))
	v.SetDefault("retry.delay", "1s")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	defaultPath := `./config`

	if configPath == "" {
		configPath = defaultPath
	}
	v.AddConfigPath(configPath)

	files, _ := ioutil.ReadDir(configPath)
	index := 0
	for _, file := range files {
		if filepath.Ext("./"+file.Name()) != ".yaml" && filepath.Ext("./"+file.Name()) != ".yml" {
			continue
		}

		v.SetConfigName(file.Name())
		var err error
		if index == 0 {
			err = v.ReadInConfig()
		} else {
			err = v.MergeInConfig()
		}
		if err == nil {
			index++
		}
	}

	if err = v.Unmarshal(&config); err != nil {
		return
	}

	return
}
