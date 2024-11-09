package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"

	"github.com/jotadrilo/cookify/internal/logger"
)

var ErrConfigDirBlank = fmt.Errorf("configuration directory is blank")

type Config struct {
	Server   Server   `json:"server" yaml:"server" mapstructure:"server"`
	Database Database `json:"database" yaml:"database" mapstructure:"database"`
}

type Server struct {
	Address string     `json:"address" yaml:"address" mapstructure:"address"`
	Gin     *ServerGin `json:"gin" yaml:"gin" mapstructure:"gin"`
}

type ServerGin struct {
	Mode string `json:"mode" yaml:"mode" mapstructure:"mode"`
}

type Database struct {
	Postgres   *DatabasePg `json:"postgres,omitempty" yaml:"postgres,omitempty" mapstructure:"postgres"`
	FileSystem *DatabaseFs `json:"fs,omitempty" yaml:"fs,omitempty" mapstructure:"fs"`
}

type DatabasePg struct {
	Host       string `json:"host" yaml:"host" mapstructure:"host"`
	Port       uint32 `json:"port" yaml:"port" mapstructure:"port"`
	User       string `json:"user" yaml:"user" mapstructure:"user"`
	Pass       string `json:"pass" yaml:"pass" mapstructure:"pass"`
	Database   string `json:"database" yaml:"database" mapstructure:"database"`
	Insecure   bool   `json:"insecure,omitempty" yaml:"insecure,omitempty" mapstructure:"insecure"`
	BunVerbose bool   `json:"bun_verbose,omitempty" yaml:"bun_verbose,omitempty" mapstructure:"bun_verbose"`
}

type DatabaseFs struct {
	Root string `json:"root" yaml:"root" mapstructure:"root"`
}

func Load() error {
	if configDir == "" {
		return fmt.Errorf("cannot load configuration: %w", ErrConfigDirBlank)
	}

	logger.Infof("Reading configuration from %q", configDir)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cannot load configuration: %w", err)
	}

	var cfg = &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("cannot load configuration: %w", err)
	}

	SetDefault(cfg)

	return nil
}

var (
	config *Config
	mu     sync.Mutex
)

func Default() *Config {
	mu.Lock()
	defer mu.Unlock()
	return config
}

func SetDefault(cfg *Config) {
	mu.Lock()
	defer mu.Unlock()
	config = cfg
}
