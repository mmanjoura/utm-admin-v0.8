package utilities

import (
	"log"
	"os"
)

func EnvStr(name string) string {
	s := os.Getenv(name)
	if s == "" {
		log.Fatal("empty config env var:", name)
	}
	return s
}
