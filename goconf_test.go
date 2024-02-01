package goconf

import (
	"testing"
)

func TestSplitLine(t *testing.T) {
	conf := NewGoConf()
	str := "Test = 12"
	key, value := conf.splitLine(str)
	if key != "Test" {
		t.Errorf("Key got \"%s\" want \"%s\"", key, "Test")
	}
	if value != "12" {
		t.Errorf("Value got \"%s\" want \"%s\"", value, "12")
	}
}
