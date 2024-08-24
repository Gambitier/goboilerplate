package config

import (
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/spf13/viper"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	registerCustomValidations(validate)

	// Set up the English translator
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)
	registerCustomTranslations(validate, trans)
}

func (c *Conf) validateConfig() []ConfigValidationError {
	var errors []ConfigValidationError

	if err := validate.Struct(c); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			field := e.Field()
			message := e.Translate(trans)
			errors = append(errors, ConfigValidationError{Field: field, Message: message})
		}
	}

	return errors
}

func loadConfig(path string) (*Conf, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json") // required if the config file does not have the extension in the name
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Conf
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	errs := cfg.validateConfig()
	if len(errs) > 0 {
		logValidationErrors(errs)
		return nil, fmt.Errorf("config validation failed")
	}

	return &cfg, nil
}

func logValidationErrors(errors []ConfigValidationError) {
	for _, err := range errors {
		log.Printf("Validation error -> '%s': %s", err.Field, err.Message)
	}
}
