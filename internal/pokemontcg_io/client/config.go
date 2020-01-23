package client

type Config struct {
	apiURL string
}

var DefaultConfig = Config{
	apiURL: defaultAPIURL,
}

func NewConfig(apiURL *string) *Config {
	c := &Config{}
	if apiURL != nil {
		c.apiURL = *apiURL
	}
}

