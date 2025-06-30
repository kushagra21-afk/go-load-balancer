package main

import (
	"os"
	"time"
	"gopkg.in/yaml.v3"
)
type HealthCheckConfig struct {
	IntervalSeconds int    `yaml:"interval_seconds"`
	Path            string `yaml:"path"`
}

type Config struct {
	Listen      string             `yaml:"listen"`
	Backends    []string           `yaml:"backends"`
	HealthCheck HealthCheckConfig  `yaml:"health_check"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (h *HealthCheckConfig) IntervalDuration() time.Duration {
	return time.Duration(h.IntervalSeconds) * time.Second
}