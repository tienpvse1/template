package common

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct{}

func (env Env) TryLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func (env Env) Get(key string) string {
	return os.Getenv(key)
}
