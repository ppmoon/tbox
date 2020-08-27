package tbox

import (
	"bytes"
	"errors"
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

// divide equally int slice
func DivideEquallyIntSlice(array []int, num int) (array2 [][]int, err error) {
	arrayLen := len(array)
	if arrayLen < num {
		err = errors.New("num must more than the arrayLen")
		return
	}
	subLen := 0
	// 每组均分大小
	if arrayLen%num == 0 {
		subLen = arrayLen / num
	} else {
		subLen = arrayLen/num + 1
	}
	for i := 0; i < num; i++ {
		start := subLen * i
		end := subLen * (i + 1)
		if i == num-1 {
			array2 = append(array2, array[start:])
		} else {
			array2 = append(array2, array[start:end])
		}
	}
	return
}
