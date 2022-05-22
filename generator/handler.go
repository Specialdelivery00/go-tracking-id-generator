package generator

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type Generator interface {
	GenerateTrackingId(stringInput string) (trackingId string)
}

type generatorSHA512 struct {
	masterKey string
	length    int
}

type generatorSHA256 struct {
	masterKey string
	length    int
}

// GenerateTrackingId is used to generate a N-length tracking ID using SHA-512 hash algo by accepting a string input.
func (generator *generatorSHA512) GenerateTrackingId(stringInput string) (trackingId string) {
	data := generator.masterKey + stringInput
	hash := sha512.New()
	hash.Write([]byte(data))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)[:generator.length]
}

// GenerateTrackingId is used to generate a N-length tracking ID using SHA-256 hash algo by accepting a string input.
func (generator *generatorSHA256) GenerateTrackingId(stringInput string) (trackingId string) {
	data := generator.masterKey + stringInput
	hash := sha256.New()
	hash.Write([]byte(data))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)[:generator.length]
}
