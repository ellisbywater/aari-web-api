package webapp

import (
	"embed"
	"io/fs"
	"path"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	embedFS      embed.FS
	unwrapFSOnce sync.Once
	unwrappedFS  fs.FS
)

func FS() fs.FS {
	unwrapFSOnce.Do(func() {
		fsys, err := fs.Sub(embedFS, "embed")
		if err != nil {
			panic(err)
		}
		unwrappedFS = fsys
	})
	return unwrappedFS
}

type Config struct {
	Service   string
	Env       string
	Debug     bool   `yaml:"debug"`
	SecretKey string `yaml:"secrety_key"`

	PGX struct {
		DSN          string `yaml:"dsn"`
		maxIdleConns int    `yaml:"maxIdleConns"`
		maxOpenConns int    `yaml:"maxOpenConns"`
		maxIdleTime  string `yaml:"maxIdleTime"`
	} `yaml:"dsn"`
}

func ReadConfig(fsys fs.FS, service, env string) (*Config, error) {
	b, err := fs.ReadFile(fsys, path.Join("config", env+".yaml"))
	if err != nil {
		return nil, err
	}
	cfg := new(Config)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}
	cfg.Service = service
	cfg.Env = env

	return cfg, nil
}
