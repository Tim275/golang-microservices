package common

import "os"

// EnvString liest eine Umgebungsvariable aus und gibt den Fallback-Wert zurück, falls die Variable nicht gesetzt ist.
func EnvString(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
