package config

import (
	log "api-bff-golang/infraestructure/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"io/ioutil"
	"os"
	"time"
)

const DURATION = time.Second * 5

var validate *validator.Validate

func LoadSettings() {
	environment := os.Getenv("ENVIRONMENT")
	log.Info("[config.go] LoadSettings | environment:")
	log.Info(environment)
	if environment == "local" || environment == "" {
		log.Info("[config.go] LoadSettings | Environment is local or empty.")
		err := getLocalConfig()
		if err != nil {
			log.Info("[config.go] LoadSettings | Se ha capturado error")
			log.WithError(err).Fatal("no config found")
		}
		log.Info("[config.go] LoadSettings | Finish Susccesfully")
		return
	}
	if len(environment) == 0 {
		log.Fatal("[config.go] LoadSettings | Not environment to run")
	}
	getRemoteConfig()
}

func getRemoteConfig() {
	log.Info("[config.go] getRemoteConfig | Init")
	pathSecret := "/etc/secrets/"

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(pathSecret)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	return
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetArray(key string) []string {
	return viper.GetStringSlice(key)
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func getLocalConfig() error {
	log.Info("[config.go] getLocalConfig | Init")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./shared/utils/config/")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	validateConfig()
	return nil
}

func validateConfig() error {

	file, _ := ioutil.ReadFile("./shared/utils/config/config.json")
	config := Config{}

	_ = json.Unmarshal([]byte(file), &config)
	validate = validator.New()
	e := validate.Struct(&config)
	if e != nil {
		log.Fatal("Error in config file with message  %s", e.Error())
	}
	return nil

}
