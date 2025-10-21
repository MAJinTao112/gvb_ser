package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

func InitConf() {
	const ConfigFile = "settings.yaml"

	c := &config.Config{}

	global.Config = c
	yamlConf, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		panic(fmt.Errorf("get yamlConf error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("yamlConf Unmarshal err:%v", err)
	}
	log.Println("config yamlFile load success")
	//fmt.Println(c)
	//fmt.Println(*c)
	//fmt.Printf("%+v\n", c)

}
func SetYaml() error {
	yamlConf, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	err = ioutil.WriteFile("settings.yaml", yamlConf, 0777)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Info("yaml file update success")
	return nil
}
