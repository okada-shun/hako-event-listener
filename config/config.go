package config

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Ethereum struct {
		ContractAddress string `yaml:"contractaddress"`
		NetworkURL      string `yaml:"networkurl"`
	}
	Mysql struct {
		Password string `yaml:"password"`
		User     string `yaml:"user"`
	}
}

func loadConfigForYaml() (*Config, error) {
	f, err := os.Open("/home/okada_shun/hako-server/config.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

// create new config instance
func NewConfig() (*Config, error) {
	cfg, err := loadConfigForYaml()
	if err != nil {
		return nil, err
	}
	return &Config{
		Ethereum: cfg.Ethereum,
		Mysql:    cfg.Mysql,
	}, nil
}
