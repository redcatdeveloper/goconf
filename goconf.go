package goconf

type goConf struct {
  Filename string `config.ini`
}

func newGoConf(filename) *goConf {
  conf := new &goConf {
    Filename: filename
  }
  return conf
}
