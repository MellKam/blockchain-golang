package converter

import (
	"bytes"
	"encoding/binary"
	"log"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func IntegerToByteSlice[T Integer](num T) []byte {
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, int64(num))
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
