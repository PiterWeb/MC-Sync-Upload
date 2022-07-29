package main

import (
	"errors"
	"github.com/joho/godotenv"
)

func getEnv(key string) (string, error) {

	if key == "" {
		return "", errors.New("Key is empty")
	}

	env, err := godotenv.Unmarshal(string(envData))

	if err != nil {
		return "", err
	}

	return env[key], nil

}
