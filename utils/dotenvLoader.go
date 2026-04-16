package utils

import "github.com/joho/godotenv"

func DotenvLoader() {
	dotenvErr := godotenv.Load()

	if dotenvErr != nil {
		panic("Failed to load env")
	}
}