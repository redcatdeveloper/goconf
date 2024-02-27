package goconf

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	conf := NewGoConf()
	status, err := conf.parseFile("test/config1.ini")
	if status {
		value := conf.Get("test")
		if value != "Hello World" {
			t.Errorf("Value got \"%s\" want \"%s\"", value, "Hello World")
		}
	} else {
		t.Errorf("Error read file: %s", err)
	}
}

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
