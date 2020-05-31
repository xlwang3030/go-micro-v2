package main

import (
	"fmt"
	"github.com/go-acme/lego/v3/log"
	"github.com/micro/go-micro/v2/config"
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
	configFile:="app.json"
	err:=config.LoadFile(configFile)
	if err!=nil{
		log.Fatal(err)
	}
	jtConfig:=&JtConfig{}
	err=config.Get("jtthink").Scan(jtConfig)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(jtConfig)

}