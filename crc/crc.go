package crc

import (
	"fmt"

	"github.com/howeyc/crc16"
)

func GenCRC(value string) string {
	data := []byte(value)
	crc1 := crc16.ChecksumCCITTFalse(data)
	fb := fmt.Sprintf("%02X", byte(crc1>>8))
	sb := fmt.Sprintf("%02X", byte(crc1&0xFF))
	return fb + sb
}
