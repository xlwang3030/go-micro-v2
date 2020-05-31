package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-acme/lego/v3/log"
	"jtthink/src/Config"
	"time"
)

func WaitForReady() error  {
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()
  for{
  	select {
  	     case <-ctx.Done():
  	     	return fmt.Errorf("init config error")
	default:
		if Config.JConfig.Data.Mysql!=nil{
			return nil
		}
	}
  }
}

func main(){

	Config.InitConfig()
	fmt.Println(Config.JConfig.Data.Mysql)
	err:=WaitForReady()
	if err!=nil{
		log.Fatal(err)
	}
	r:=gin.New()
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Config.JConfig.Data.Mysql})
	})
	r.Run(":8080")



}