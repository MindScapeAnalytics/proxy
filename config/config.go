package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	Server                      Server
	Logger                      Logger
	AuthenticationService       AuthenticationService
	VisualRepresentationService VisualRepresentationService
	PsychologyTestingService    PsychologyTestingService
}

type Server struct {
	IP                          string `validate:"required"`
	Port                        string `validate:"required"`
	ShowUnknownErrorsInResponse bool   `validate:"required"`
	SecretJWT                   string `validate:"required"`
	Instance                    string `validate:"required"`
}

type AuthenticationService struct {
	IP   string `validate:"required"`
	Port string `validate:"required"`
}

type VisualRepresentationService struct {
	IP   string `validate:"required"`
	Port string `validate:"required"`
}

type PsychologyTestingService struct {
	IP   string `validate:"required"`
	Port string `validate:"required"`
}

type Logger struct {
	Level zerolog.Level `validate:"required"`
	File  string
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}
	return
}
