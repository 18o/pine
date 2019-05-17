package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path/filepath"
	"pine/config"
	"pine/pkg/fs"
)

var (
	app        string
	configPath string
	err        error
	r          = gin.Default()
)

func init() {
	flag.StringVar(&app, "a", "pine", "specify server name")
	flag.Parse()
	configPath = "/etc/pine/" + app + ".toml"
	if !fs.IsExists(configPath) {
		configPath, err = filepath.Abs("config/app.toml")
		if err != nil {
			log.Error(err)
			return
		}
		if !fs.IsExists(configPath) {
			log.Panic("配置文件不存在: %s", configPath)
		}
	}

	viper.SetConfigType("toml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("读取配置文件错误: %s", err.Error())
		return
	}
	log.Println("using configure file " + configPath)
	var ginModal = viper.GetString("global.gin-modal")
	gin.SetMode(ginModal)

}

func main() {
	config.Route(r)
	err := r.Run(viper.GetString("global.addr"))
	if err != nil {
		log.Panic(err)
	}
}
