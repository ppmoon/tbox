package tbox

import (
	"bytes"
	"io"
)

// io.reader to []byte
func IoReaderToByte(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	return buf.Bytes(), err
}

// io.reader to string
func IoReaderToString(stream io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	return buf.String(), err
}
