package generator

import (
	"crypto/sha1"
	"fmt")

type Generator interface{
	GenerateTrackingId(stringInput string)(trackingId string)
}

type generator struct{
	masterKey string
	length int
}

func newGenerator(masterKey string, length int) Generator{
	return &generator{masterKey: masterKey,length: length}
}

func (generator *generator) GenerateTrackingId(stringInput string) (trackingId string){
	data := generator.masterKey + stringInput
	hash := sha1.New()
	hash.Write([]byte(data))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)[:generator.length]
}

