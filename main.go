package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func GetUUIDv7() ([16]byte, error) {
	// random bytes
	var value [16]byte
	_, err := rand.Read(value[:])
	if err != nil {
		return value, err
	}

	// current timestamp in ms
	timestamp := big.NewInt(time.Now().UnixMilli())

	// timestamp
	timestamp.FillBytes(value[0:6])

	// version and variant
	value[6] = (value[6] & 0x0F) | 0x70
	value[8] = (value[8] & 0x3F) | 0x80

	return value, nil
}

func GetUUIDv7String() (string, error) {
	uuidVal, err := GetUUIDv7()
	if err != nil {
		return "", err
	}

	return convertUuidToString(uuidVal), nil
}

func convertUuidToString(uuidVal [16]byte) string {
	// format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	result := fmt.Sprintf("%x-%x-%x-%x-%x", uuidVal[:4], uuidVal[4:6], uuidVal[6:8], uuidVal[8:10], uuidVal[10:])
	return result
}
