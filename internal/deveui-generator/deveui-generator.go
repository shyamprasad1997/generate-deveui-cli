package deveuigenerator

import (
	"crypto/rand"
	"math/big"

	log "github.com/sirupsen/logrus"
)

const allowedChars = "ABCDEF0123456789"

func Generate() []string {
	uniqueCheckMap := make(map[string]bool)
	log.Print("starting generating deveuis")
	var devEuis []string
	for i := 0; i < 100; i++ {
		devEui, err := generateHexString(16)
		if err != nil {
			log.Fatal("failed to generate hexstring, err in generateHexString(), err: ", err.Error())
		}
		if uniqueCheckMap[devEui[len(devEui)-5:]] {
			i = i - 1
			continue
		}
		uniqueCheckMap[devEui[len(devEui)-5:]] = true
		devEuis = append(devEuis, devEui)
	}
	log.Print("completed generating deveuis")
	return devEuis
}

func generateHexString(length int) (string, error) {
	max := big.NewInt(int64(len(allowedChars)))
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b[i] = allowedChars[n.Int64()]
	}
	return string(b), nil
}
