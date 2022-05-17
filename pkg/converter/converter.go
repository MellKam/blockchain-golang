package converter

import (
	"bytes"
	"encoding/binary"
	"log"
)

func UintToByteSlice(num uint) []byte {
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, int64(num))
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
