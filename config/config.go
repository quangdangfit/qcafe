package config

type Schema struct {
	MongoDB struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Database string `mapstructure:"database"`
		User     string `mapstructure:"user"`
		Pass     string `mapstructure:"pass"`
		Env      string `mapstructure:"env"`
	}
}
