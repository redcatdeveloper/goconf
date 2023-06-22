package goconf

type goConf struct {
  Filename string `config.ini`
}

func newGoConf(filename string) *goConf {
  conf := &goConf {
    Filename: filename,
  }
  return conf
}
