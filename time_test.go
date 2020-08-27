package tbox_test

import (
	"github.com/ppmoon/tbox"
	"testing"
)

func TestFromNowToTomorrowMidnight(t *testing.T) {
	t.Log(tbox.FromNowToTomorrowMidnight())
}
