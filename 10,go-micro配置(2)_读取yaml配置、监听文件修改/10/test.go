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

	configFile:="app.yaml"

	err:=config.LoadFile(configFile)
	if err!=nil{
		log.Fatal(err)
	}
	jtconfig:=&JtConfig{}
	go func() {
		for {
			w,_:=config.Watch("jtthink")
			v,_:=w.Next()//卡主

		}
	}()

	v.Scan(jtconfig)
	fmt.Println(jtconfig)



}