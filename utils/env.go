package utils

import "os"

func GetEnv(EnvKey string) string {
	var EnvKeyValue = os.Getenv(EnvKey)

	return EnvKeyValue
}
