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
	trackingId := trackIdGen.GenerateTrackingId(sampleStringInput)
	fmt.Println(trackingId)
}
