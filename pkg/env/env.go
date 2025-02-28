package env

import (
	"os"
	"strconv"
	"time"
)

// GetString récupère une variable d'environnement sous forme de string
func GetString(key, fallback string) string {
	if val, found := os.LookupEnv(key); found {
		return val
	}
	return fallback
}

// GetInt récupère une variable d'environnement sous forme d'entier
func GetInt(key string, fallback int) (int, error) {
	val, found := os.LookupEnv(key)
	if !found {
		return fallback, nil
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback, err
	}
	return intVal, nil
}

// GetBool récupère une variable d'environnement sous forme de booléen
func GetBool(key string, fallback bool) (bool, error) {
	val, found := os.LookupEnv(key)
	if !found {
		return fallback, nil
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return fallback, err
	}
	return boolVal, nil
}

// GetFloat récupère une variable d'environnement sous forme de float64
func GetFloat(key string, fallback float64) (float64, error) {
	val, found := os.LookupEnv(key)
	if !found {
		return fallback, nil
	}

	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fallback, err
	}
	return floatVal, nil
}

// GetDuration récupère une variable d'environnement sous forme de durée (time.Duration)
func GetDuration(key string, fallback time.Duration) (time.Duration, error) {
	val, found := os.LookupEnv(key)
	if !found {
		return fallback, nil
	}

	durationVal, err := time.ParseDuration(val)
	if err != nil {
		return fallback, err
	}
	return durationVal, nil
}
