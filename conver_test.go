package tbox_test

import (
	"reflect"
	"tbox"
	"testing"
)

func Test_IntSliceToStringSlice(t *testing.T) {
	a := []int{1, 2, 3}
	b := tbox.IntSliceToStringSlice(a)
	c := []string{"1", "2", "3"}
	if !reflect.DeepEqual(b, c) {
		t.Error("int slice to string slice error")
		return
	}
}

func TestStringSliceToIntSlice(t *testing.T) {
	a := []string{"1", "2", "3"}
	b, err := tbox.StringSliceToIntSlice(a)
	if err != nil {
		t.Error(err)
		return
	}
	c := []int{1, 2, 3}
	if !reflect.DeepEqual(b, c) {
		t.Error("string slice to int slice error")
		return
	}
}
