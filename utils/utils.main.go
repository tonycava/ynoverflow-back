package utils

import (
	"os"
	"strconv"
)

func GetEnv(name string) string {
	env, isValid := os.LookupEnv(name)
	if !isValid {
		panic("Environment variable " + name + " not found")
	}
	return env
}

func ParseInt(nbr string) int {
	intNBr, _ := strconv.Atoi(nbr)
	return intNBr
}
