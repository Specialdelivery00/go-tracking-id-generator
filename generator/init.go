package generator

// NewSHA512Generator is used to generate a new generator
func NewSHA512Generator(masterKey string, length int) Generator {
	return &generatorSHA512{masterKey: masterKey, length: length}
}

// NewSHA256Generator is used to generate a new generator
func NewSHA256Generator(masterKey string, length int) Generator {
	return &generatorSHA256{masterKey: masterKey, length: length}
}
