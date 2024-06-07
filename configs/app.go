package configs

type appConfig struct {
	Port    string `envconfig:"APP_PORT" default:"8081"`
	ENV     string `envconfig:"APP_ENV" default:"local"`
	Version string `envconfig:"APP_VERSION" default:"v0.0.0"`
	APIKey  string `envconfig:"APP_API_KEY" default:"this-is-api-key"`
	Prefix  string `envconfig:"APP_PREFIX" default:"app"`
	Name    string `envconfig:"APP_NAME" default:"application"`
	Log     string `envconfig:"APP_PATH_LOG" default:"logs"`
}
