package tbox

import (
	"bytes"
	"io"
	"strconv"
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

// []int to []string
func IntSliceToStringSlice(intSlice []int) (stringSlice []string) {
	for _, i := range intSlice {
		stringSlice = append(stringSlice, strconv.Itoa(i))
	}
	return
}

// []string to []int
func StringSliceToIntSlice(stringSlice []string) (intSlice []int, err error) {
	for _, i := range stringSlice {
		var stringInt int
		stringInt, err = strconv.Atoi(i)
		if err != nil {
			return
		}
		intSlice = append(intSlice, stringInt)
	}
	return
}
