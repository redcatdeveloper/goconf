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

func getTestConfig(name string) *goConf {
	conf := NewGoConf()
	status := conf.LoadFile("test/" + name)
	if status {
		return conf
	}
	return nil
}

func TestGet(t *testing.T) {
	testfile := "config1.ini"
	conf := getTestConfig(testfile)
	if conf == nil {
		t.Fatalf("Can't load test file %s", testfile)
	}
	value := conf.Get("test")
	if value != "Hello World" {
		t.Errorf("Value got \"%s\" want \"%s\"", value, "Hello World")
	}
}

func TestGetInt(t *testing.T) {
	testfile := "config1.ini"
	key := "testInt"
	conf := getTestConfig(testfile)
	if conf == nil {
		t.Fatalf("Can't load test file %s", testfile)
	}
	value := conf.Get(key)
	if value != "10" {
		t.Errorf("Value got \"%s\" want \"%s\"", value, "10")
	}
	value1, err := conf.GetInt(key)
	if err != nil {
		t.Fatalf("Can't get Int value from \"%s\": %v", key, err)
	}
	if (value1 != 10)	{
		t.Errorf("Value got \"%d\" want \"%d\"", value1, 10)
	}
}
