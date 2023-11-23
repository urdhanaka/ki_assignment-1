package utils

import "os"

func GetEnv(EnvKey string) string {
	EnvKeyValue := os.Getenv(EnvKey)

	return EnvKeyValue
}
