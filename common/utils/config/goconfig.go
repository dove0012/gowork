package config

import (
	"github.com/Unknwon/goconfig"
	"gowork/common/utils/log"
)

type Cfg struct {
	c *goconfig.ConfigFile
}

var C = &Cfg{}

func init()  {
	cfg, err := goconfig.LoadConfigFile("config/config.ini")
	log.Error2Exit(err, "goconfig.LoadConfigFile error")
	C.c = cfg
}

func NewCfg(fileName string) *Cfg {
	if fileName == "config.ini" {
		return C
	}
	cfg, err := goconfig.LoadConfigFile("config/" + fileName)
	log.Error2Exit(err, "goconfig.LoadConfigFile error")
	return &Cfg{c: cfg}
}

func (cfg *Cfg) GetString(section string, key string) string {
	v, err := cfg.c.GetValue(section, key)
	log.Error2Exit(err, "goconfig.GetValue error")
	return v
}

func (cfg *Cfg) GetStringByDefault(key string) string {
	v, err := cfg.c.GetValue(goconfig.DEFAULT_SECTION, key)
	log.Error2Exit(err, "goconfig.GetValue error")
	return v
}

func (cfg *Cfg) GetInt(section string, key string) int64 {
	v, err := cfg.c.Int64(section, key)
	log.Error2Exit(err, "goconfig.Int64 error")
	return v
}

func (cfg *Cfg) GetIntByDefault(key string) int64 {
	v, err := cfg.c.Int64(goconfig.DEFAULT_SECTION, key)
	log.Error2Exit(err, "goconfig.Int64 error")
	return v
}

func (cfg *Cfg) GetFloat(section string, key string) float64 {
	v, err := cfg.c.Float64(section, key)
	log.Error2Exit(err, "goconfig.Float64 error")
	return v
}

func (cfg *Cfg) GetFloatByDefault(key string) float64 {
	v, err := cfg.c.Float64(goconfig.DEFAULT_SECTION, key)
	log.Error2Exit(err, "goconfig.Float64 error")
	return v
}
