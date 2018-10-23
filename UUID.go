package utils

import (
	"github.com/satori/go.uuid"
	"github.com/gobasis/log"
	"encoding/hex"
)

type uuidType struct {}

var UUID uuidType

func (u uuidType) Get() string {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Error("uuid create error", "error", err)
		return ""
	}
	buf := make([]byte, 32)
	hex.Encode(buf[0:8], uid[0:4])
	hex.Encode(buf[8:12], uid[4:6])
	hex.Encode(buf[12:16], uid[6:8])
	hex.Encode(buf[16:20], uid[8:10])
	hex.Encode(buf[20:], uid[10:])
	return string(buf)
}