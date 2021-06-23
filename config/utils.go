package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Load :
// 設定ファイルを読み込む
func (c *Config) Load(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Println("Config successfully loaded! ", *c)

	return nil
}
