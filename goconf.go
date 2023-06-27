package goconf

import (
  "os"
  "bufio"
  "strings"
  "log"
)

type inidata map[string]string

type goConf struct {
  Filename string `config.ini`
  data inidata
}

func NewGoConf(filename string) (*goConf, error) {
  conf := &goConf {
    Filename: filename,
    data: inidata{},
  }
  _, err := conf.parseFile()
  if err != nil {
    return nil, err
  }
  return conf, nil
}

func (conf *goConf) Get(param string) string {
  value, ok := conf.data[param]
  if ok {
    return value
  }
  return ""
}

func (conf *goConf) parseFile() (bool, error) {
  f, err := os.OpenFile(conf.Filename, os.O_RDONLY, 0666)
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

func (conf *goConf) parseLine(line string) {
  parts := strings.Split(strings.Trim(line, "\r\n\t "), "=")
  if len(parts) == 2 {
    key, value := strings.Trim(parts[0], "\r\n\t "), strings.Trim(parts[1], "\r\n\t ")
    if key != "" {
      conf.data[key] = value
    }
  }
}
