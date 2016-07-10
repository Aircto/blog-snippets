package envron

import (
	"log"
	"os"
)

func MustEnv(key string) (value string) {
	if value = os.Getenv(key); value == "" {
		log.Fatalf("ENV %q is not set.", key)
	}
	return value
}
