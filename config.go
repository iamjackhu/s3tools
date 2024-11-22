package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type S3Endpoint struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Region    string `yaml:"region"`
}

type S3ToolConfig struct {
	S3Endpoint    S3Endpoint `yaml:"s3Endpoint"`
	BucketPrefile string     `yaml:"bucketPrefile"`
	Concurrent    int        `yaml:"concurrent"`
	Total         int        `yaml:"total"`
}

// GetConfig will read the configuratio from yaml file
func GetConfig(configFile string) (*S3ToolConfig, error) {
	f, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var cfg S3ToolConfig
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	setDefaultValue(&cfg)

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setDefaultValue(cfg *S3ToolConfig) {
	if cfg.Concurrent == 0 {
		cfg.Concurrent = 1
	}

	if cfg.Total == 0 {
		cfg.Total = 1
	}

	if cfg.BucketPrefile == "" {
		cfg.BucketPrefile = "ppbucket"
	}

	if cfg.S3Endpoint.Region == "" {
		cfg.S3Endpoint.Region = "default"
	}
}

func validate(cfg *S3ToolConfig) error {

	if cfg.Concurrent > cfg.Total {
		return fmt.Errorf("the concurrent %d should not over total %d", cfg.Concurrent, cfg.Total)
	}

	if cfg.S3Endpoint.Endpoint == "" || cfg.S3Endpoint.AccessKey == "" || cfg.S3Endpoint.SecretKey == "" {
		return fmt.Errorf("the s3 endpoint info should not be empty")
	}

	return nil
}
