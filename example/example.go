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
	trackIdGen := generator.NewGenerator(masterKey, length)
	trackingId := trackIdGen.GenerateSHA512TrackingId(sampleStringInput)
	fmt.Println("SHA-512:" + trackingId)
	trackingId = trackIdGen.GenerateSHA256TrackingId(sampleStringInput)
	fmt.Println("SHA-256:" + trackingId)
}
