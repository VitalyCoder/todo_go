package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envcnofig:"ADDR" required:"true"`
	ShutdownTimeout time.Duration `envcnofig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("Process envconfig: %w", err)
	}

	return config, nil
}

func ConfigMust() Config {
	config, err := NewConfig()

	if err != nil {
		err := fmt.Errorf("get HTTP server config: %w", err)
		panic(err)
	}

	return config
}
