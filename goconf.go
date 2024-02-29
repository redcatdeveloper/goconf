package goconf

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type inidata map[string]string

type goConf struct {
	data inidata
}

func NewGoConf() *goConf {
	conf := &goConf{
		data: inidata{},
	}
	return conf
}

func (conf *goConf) LoadFile(filename string) bool {
	_, err := conf.parseFile(filename)
	if err != nil {
		return false
	}
	return true
}

/**
Get string value from config by name
**/
func (conf *goConf) Get(param string) (val string, err error) {
	val, ok := conf.data[param]
	if !ok {
		err = errors.New("Not found key: " + param)
		log.Printf("Not found key \"%s\"", param)
	}
	return
}

/**
Get int value from config by name
**/
func (conf *goConf) GetInt(param string) (val int, err error) {
	var value string
	value, err = conf.Get(param)
	if err != nil {
		log.Print(err)
		return
	}
	if value == "" {
		err = errors.New("Empty value for " + param)
		log.Print(err)
	}
	val, err = strconv.Atoi(value)
	return
}

func (conf *goConf) GetArray(param string, options ...string) (val []string, err error) {
	var value, sep string
	sep = ";"
	value, err = conf.Get(param)
	if err != nil {
		log.Print(err)
		return
	}
	if value == "" {
		return
	}
	if len(options) > 0 {
		sep = options[0]
	}
	val = strings.Split(value, sep)
	return
}

func (conf *goConf) parseFile(filename string) (bool, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Print(err)
		return false, err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		conf.parseLine(line)
	}
	return true, nil
}

func (conf *goConf) splitLine(line string) (key string, value string) {
	parts := strings.Split(strings.Trim(line, "\r\n\t "), "=")
	if len(parts) == 2 {
		key, value = strings.Trim(parts[0], "\r\n\t "), strings.Trim(parts[1], "\r\n\t ")
	}
	return
}

func (conf *goConf) parseLine(line string) bool {
	key, value := conf.splitLine(line)
	if key != "" {
		conf.data[key] = value
		return true
	}
	return false
}
