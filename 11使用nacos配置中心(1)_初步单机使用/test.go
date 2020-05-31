package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"time"
)

type JtConfig struct {
	Db struct{
		Ip string `json:"ip"`
		Port int `json:"port"`
	} `json:"db"`
	Redis struct{
		Ip string `json:"ip"`
		Port int `json:"port"`
	} `json:"redis"`
}
func main(){
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "localhost",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err!=nil{
		log.Fatal(err)
	}
	 err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "jt-sysconfig",
		Group:  "jtthink",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(data)
		},
	 })
	if err!=nil{
		log.Fatal(err)
	}
	for{
		time.Sleep(time.Second*1)
	}


}