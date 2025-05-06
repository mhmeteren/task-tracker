package util

import (
	"log"

	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateKey(length int) string {
	key, err := nanoid.New(length)
	if err != nil {
		log.Fatal("Key generation error:", err)
	}
	return key
}
