package tbox_test

import (
	"tbox"
	"testing"
)

func TestFromNowToTomorrowMidnight(t *testing.T) {
	t.Log(tbox.FromNowToTomorrowMidnight())
}
