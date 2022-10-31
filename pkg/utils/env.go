package utils

import "os"

func GetEnvOrDefault(envName, defaultValue string) string {
	value, exist := os.LookupEnv(envName)
	if exist {
		return value
	}
	return defaultValue
}

func MustGetEnv(envName string) string {
	value, exist := os.LookupEnv(envName)
	if exist {
		return value
	}
	panic("env [" + envName + "] not exist")
}
