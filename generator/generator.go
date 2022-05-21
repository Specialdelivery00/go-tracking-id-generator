package generator

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type Generator interface {
	GenerateSHA512TrackingId(stringInput string) (trackingId string)
	GenerateSHA256TrackingId(stringInput string) (trackingId string)
}

type generator struct {
	masterKey string
	length    int
}

// NewGenerator is used to generate a new generator
func NewGenerator(masterKey string, length int) Generator {
	return &generator{masterKey: masterKey, length: length}
}

// GenerateSHA512TrackingId is used to generate a N-length tracking ID using SHA-512 hash algo by accepting a string input.
func (generator *generator) GenerateSHA512TrackingId(stringInput string) (trackingId string) {
	data := generator.masterKey + stringInput
	hash := sha512.New()
	hash.Write([]byte(data))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)[:generator.length]
}

// GenerateSHA256TrackingId is used to generate a N-length tracking ID using SHA-256 hash algo by accepting a string input.
func (generator *generator) GenerateSHA256TrackingId(stringInput string) (trackingId string) {
	data := generator.masterKey + stringInput
	hash := sha256.New()
	hash.Write([]byte(data))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)[:generator.length]
}
