package g

import (
	"fmt"
	"github.com/toolkits/file"
	"gopkg.in/yaml.v2"
	"log"
	"sync"
)

type Config struct {
	MysqlConfig MysqlConfig `json:"mysqlConfig" yaml:"store"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
}

var (
	Cfg        *Config
	configLock = new(sync.RWMutex)
)

func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("configuration file %s is nonexistent", cfg)
	}

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}

	var c Config
	err = yaml.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	configLock.Lock()
	defer configLock.Unlock()
	Cfg = &c
	log.Println(Cfg)

	log.Println("load configuration file", cfg, "successfully")
	return nil
}
