package main

import (
	"fmt"
	"jtthink/src/Config"
	"time"
)

func main(){

	Config.InitConfig()
	fmt.Println(Config.JConfig.Data.Mysql)

	for{
		time.Sleep(time.Second*1)
		fmt.Println(Config.JConfig.Data.Mysql)
	}


}