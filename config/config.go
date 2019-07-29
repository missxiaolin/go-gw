package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type configure struct {
	Produce Produce			`yaml:"produce"`
	QiNiu	QiNiuConfig		`yaml:"qiniu"`
}

type QiNiuConfig struct {
	Ak		string			`yaml:"ak"`		// ak
	Sk		string			`yaml:"sk"`		// sk
	Bucket	string			`yaml:"bucket"`	// bucket
}

type Produce struct {
	Port            string `yaml:"port"`              // 监听端口
}

var Cfg = &configure{}

func init() {

	var err error
	b, err := ioutil.ReadFile(CONFIG_PATH)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(b, Cfg); err != nil {
		panic(err)
	}


	if Cfg.Produce.Port == "" {
		Cfg.Produce.Port = PRODUCE_PORT
	}
}