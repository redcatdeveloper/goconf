package goconf

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	conf := NewGoConf()
	status, err := conf.parseFile("test/config1.ini")
	if status {
		value, _ := conf.Get("test")
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
	value, err := conf.Get("test")
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if value != "Hello World" {
		t.Errorf("Value got \"%s\" want \"%s\"", value, "Hello World")
	}

	_, err1 := conf.Get("test1")
	if err1 == nil {
		t.Errorf("Should by error for wrong key \"test1\"")
	}
}

func TestGetInt(t *testing.T) {
	testfile := "config1.ini"
	key := "testInt"
	conf := getTestConfig(testfile)
	if conf == nil {
		t.Fatalf("Can't load test file %s", testfile)
	}

	value, err := conf.Get(key)
	if value != "10" {
		t.Errorf("Value got \"%s\" want \"%s\"", value, "10")
	}

	value1, err1 := conf.GetInt(key)
	if err1 != nil {
		t.Errorf("Can't get Int value from \"%s\": %v", key, err)
	} else {
		if value1 != 10 {
			t.Errorf("Value got \"%d\" want \"%d\"", value1, 10)
		}
	}

	_, err2 := conf.GetInt("test2")
	if err2 == nil {
		t.Errorf("Should by error for wrong key \"test2\"")
	}
}

func TestGetArray(t *testing.T) {
	testfile := "config1.ini"
	key := "testArr"
	conf := getTestConfig(testfile)
	if conf == nil {
		t.Fatalf("Can't load test file %s", testfile)
	}
	arr, err := conf.GetArray(key)
	if err != nil {
		t.Errorf("Can't get Array value from \"%s\": %v", key, err)
	}
	if arr[0] != "a" || arr[1] != "b" || arr[2] != "c" {
		t.Errorf("Value got %v", arr)
	}
}
