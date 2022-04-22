package enviroment

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	SetupEnviroment()
}

func SetupEnviroment() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
}

func Get(i string) string {
	return os.Getenv(i)
}

func Set(key string, i string) string {
	os.Setenv(key, i)

	return os.Getenv(key)
}
