package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

var env ConfigDto

func init() {
	if env.port == "" {
		LoadEnvironmentVariable()
	}
	ConfigEnv()
}

type ConfigDto struct {
	port          string
	secret_key    string
	database_url  string
	database_name string
	debug         string
}

func ConfigEnv() {
	env = ConfigDto{
		port:          os.Getenv("PORT"),
		secret_key:    os.Getenv("SECRET_KEY"),
		database_url:  os.Getenv("MONGO_DB_URL"),
		database_name: os.Getenv("MONGO_DB_NAME"),
		debug:         os.Getenv("DEBUG"),
	}
}

func LoadEnvironmentVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env variable!")
	}
}

func accessField(key string) (string, error) {
	v := reflect.ValueOf(env)
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("Expected Struct")
	}
	_, ok := t.FieldByName(key)
	if !ok {
		return "", fmt.Errorf("property could not be found!")
	}
	f := v.FieldByName(key)
	return f.String(), nil
}

func GetEnvProperty(key string) string {
	if env.port == "" {
		ConfigEnv()
	}
	value, err := accessField(key)
	if err != nil {
		fmt.Println(err)
		return value
	}
	return value
}
