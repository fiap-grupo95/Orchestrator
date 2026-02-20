package config

import "os"

type Config struct {
	Port             string
	OSBaseURL        string
	OSAuthToken      string
	BillingBaseURL   string
	ExecutionBaseURL string
	EntityBaseURL    string
}

func Load() Config {
	return Config{
		Port:             getEnv("PORT", "8080"),
		OSBaseURL:        getEnv("OS_BASE_URL", "http://os-service:8080"),
		OSAuthToken:      getEnv("OS_AUTH_TOKEN", ""),
		BillingBaseURL:   getEnv("BILLING_BASE_URL", "http://billing-service:8080"),
		ExecutionBaseURL: getEnv("EXEC_BASE_URL", "http://execution-service:8080"),
		EntityBaseURL:    getEnv("ENTITY_BASE_URL", "http://entity-api-service:8080"),
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
