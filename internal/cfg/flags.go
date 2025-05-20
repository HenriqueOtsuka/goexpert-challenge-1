package cfg

import (
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	// App flags
	&cli.StringFlag{
		Name:        "app-env",
		Destination: &App.Environment,
		EnvVar:      "APP_ENV",
		Value:       Development,
	},
	&cli.StringFlag{
		Name:        "app-log-level",
		Destination: &App.LogLevel,
		EnvVar:      "APP_LOG_LEVEL",
		Value:       "debug",
	},

	// Free Wheather API flags
	&cli.StringFlag{
		Name:        "freeweather-api-key",
		Destination: &FreeWeather.ApiKey,
		EnvVar:      "FREEWEATHER_API_KEY",
	},
}
