package generator

import (
	"crypto/sha512"
	"fmt"
)

type Generator interface {
	GenerateTrackingId(stringInput string) (trackingId string)
}

type generator struct {
	masterKey string
	length    int
}

// NewGenerator is used to generate a new generator
func NewGenerator(masterKey string, length int) Generator {
	return &generator{masterKey: masterKey, length: length}
}

// GenerateTrackingId is used to generate a N-length tracking ID by accepting a string input.
func (generator *generator) GenerateTrackingId(stringInput string) (trackingId string) {
	data := generator.masterKey + stringInput
	hash := sha512.New()
	hash.Write([]byte(data))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)[:generator.length]
}
