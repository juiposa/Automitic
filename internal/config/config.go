package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	FFLogsApiEndpoint    string `envconfig:"AUTOMITIC_FFLOGS_API_ENDPOINT" default:"https://www.fflogs.com/api/v2/client"`
	FFLogsClientID       string `envconfig:"AUTOMITIC_FFLOGS_CLIENT_ID"`
	FFLogsClientSecret   string `envconfig:"AUTOMITIC_FFLOGS_CLIENT_SECRET"`
	FFLogsAccessTokenURL string `envconfig:"AUTOMITIC_FFLOGS_ACCESS_TOKEN_URL"`
	FightCacheLocation   string `envconfig:"AUTOMITIC_FIGHT_CACHE_LOCATION" default:"~/.automitic/cache"`
	DataLocation         string `envconfig:"AUTOMITIC_DATA_LOCATION" default:"~/.automitic/data"` //TODO expand home directory
}

func LoadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("automitic", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
