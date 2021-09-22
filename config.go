package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	env "github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Load() {
	bootstrapEnv()
	bootstrapConfig()
}

func Get(key string) string {
	var value string
	viper.UnmarshalKey(key, &value, viper.DecodeHook(decodeHook))

	return value
}

func bootstrapEnv() {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(fmt.Errorf("Fatal error loading .env"))
	}
}

func bootstrapConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func decodeHook (f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String {
		stringData := data.(string)
		if strings.HasPrefix(stringData, "${") && strings.HasSuffix(stringData, "}") {
			envVarValue := os.Getenv(strings.TrimPrefix(strings.TrimSuffix(stringData, "}"), "${"))
			if len(envVarValue) > 0 {
				return envVarValue, nil
			}
		}
	}
	return data, nil
}