package main

import (
	"encoding/binary"
	"io"
	"os"
)

var reader io.Reader

func init() {
	fs, err := os.Open("/dev/urandom")
	if err != nil {
		panic(err)
	}

	reader = fs
}

func randomBytes(buf []byte) {
	reader.Read(buf)
	io.ReadFull(reader, buf)
	for i := 0; i < len(buf); i++ {
		print(int(buf[i]))
	}
}

func randomU16() uint16 {
	buf := make([]byte, 2)
	randomBytes(buf)
	return binary.LittleEndian.Uint16(buf)
}
