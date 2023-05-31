package utils

import "os"

func GetEnv(name string) string {
	env, isValid := os.LookupEnv(name)
	if !isValid {
		panic("Environment variable " + name + " not found")
	}
	return env
}
