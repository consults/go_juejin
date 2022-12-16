package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	config *Yaml
)

func GetConfig() *Yaml {
	return config
}

type User struct {
	Cookie string `yaml:"cookie"`
	MsgId  string `yaml:"msgId"`
	Name   string `yaml:"name"`
}

type Yaml struct {
	User []*User `yaml:"user"`
	Time string  `yaml:"time"`
}

func Init() {
	configPath := "conf/config.yaml"
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败")
	}
	//fmt.Println(viper.Get("user.cookie"))
	y := new(Yaml)
	err1 := viper.Unmarshal(y)
	if err1 != nil {
		log.Fatal("配置失败")
	}
	config = y
	log.Println("读取配置文件成功")
	for _, user := range config.User {
		log.Println("=================")
		log.Println(user.Name, "===", user.MsgId)
		log.Println("=================")
	}
	time := os.Getenv("Time")
	if time != "" {
		config.Time = time
	}
}
