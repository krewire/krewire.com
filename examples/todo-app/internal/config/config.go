// Package config defines the typed application configuration
// krewire.yaml: metadata and project-level config (metadata only)
package config

import (
	"fmt"
	"os"

	rconfig "github.com/krewire/krewire/pkg/config"
	validation "github.com/krewire/krewire/pkg/validation"
)

// Metadata mirrors krewire.yaml.
type Metadata struct {
	Project Project `yaml:"project"`
	Dev     Dev     `yaml:"dev"`
}

// Project holds the project section of krewire.yaml.
type Project struct {
	Name    string `yaml:"name" validate:"required"`
	Kind    string `yaml:"kind" validate:"required"`
	Version string `yaml:"version"`
}

// Dev holds the dev section of krewire.yaml.
type Dev struct {
	Port int `yaml:"port"`
}

// Addr returns the listen address derived from env or krewire.yaml dev.port.
func (m *Metadata) Addr() string {
	if addr := os.Getenv("APP_ADDR"); addr != "" {
		return addr
	}
	if m.Dev.Port > 0 {
		return fmt.Sprintf(":%d", m.Dev.Port)
	}
	return ":8080"
}

// Env returns the environment variable for key, or fallback if unset or empty (KWL-CFGV-005).
func Env(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// LoadMetadata reads krewire.yaml from path, overlays the environment, and returns a
// validated Metadata.
func LoadMetadata(path string) (*Metadata, error) {
	cfg := &Metadata{}
	if err := rconfig.Load(path, cfg); err != nil {
		return nil, err
	}
	if err := rconfig.Override(cfg, os.LookupEnv); err != nil {
		return nil, err
	}
	if err := validation.Struct(cfg); err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}
	return cfg, nil
}
