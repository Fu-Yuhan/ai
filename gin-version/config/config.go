package config

type Config struct {
	App    `yaml:"App"`
	Redis  `yaml:"Redis"`
	Postgres  `yaml:"Postgres"`
	Log    `yaml:"Log"`
	CacheX `yaml:"CacheX"`
}
