package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	MySql DbCfg `yaml:"mysql"`
	PgSql DbCfg `yaml:"pgsql"`
}

type DbCfg struct {
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Schema   string `yaml:"schema"`
}

func LoadConf(fileName string, cfg *Conf) error {
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading config file: %s", err)
	}
	err = yaml.Unmarshal(byteValue, cfg)
	if err != nil {
		return fmt.Errorf("error during umarshal: %s", err)
	}
	return err
}

// TODO define Conf structure
// And write a load-from-file function
