package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	cqmodels "gitlab.qarea.org/jiraquality/cq-web-backend/common-lib/models"
)

// Config config
type Config struct {
	CommonConfig             cqmodels.Config
	ForcedCORS               bool   `toml:"forced_cors_headers"`
	ObserverCheckEveryMinute uint64 `toml:"observer_check_every_minute"`
	BindedVersion            string
}

// Load configuration from file
func Load(localpath, commonpath string) (*Config, error) {
	cfg := &Config{}
	if _, err := toml.DecodeFile(localpath, &cfg); err != nil {
		return nil, errors.New("error loading config:" + err.Error())
	}
	if _, err := toml.DecodeFile(commonpath, &cfg.CommonConfig); err != nil {
		return nil, errors.New("error loading comon config:" + err.Error())
	}
	return cfg, nil
}
