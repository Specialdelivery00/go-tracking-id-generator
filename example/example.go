package main

import (
	"fmt"
	"github.com/go-tracking-id-generator/generator"
)

const (
	masterKey = "example"
	length    = 16
)

func main() {
	sampleStringInput := "Sample1"
	sha512TrackIdGen := generator.NewSHA512Generator(masterKey, length)
	sha256TrackIdGen := generator.NewSHA256Generator(masterKey, length)
	trackingId := sha512TrackIdGen.GenerateTrackingId(sampleStringInput)
	fmt.Println("SHA-512:" + trackingId)
	trackingId = sha256TrackIdGen.GenerateTrackingId(sampleStringInput)
	fmt.Println("SHA-256:" + trackingId)
}
