package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"jtthink/src/Config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitForReady() error  {
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*8)
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

func CheckForReady() gin.HandlerFunc {
	return func(i *gin.Context) {
		if Config.JConfig.Data.Mysql==nil{
			i.AbortWithStatusJSON(200,gin.H{"result":"server is loading"})
		}else{
			i.Next()
		}
	}
}
func main(){
	Config.InitConfig()
	errChan:=make(chan error)
	go func() {
		err:=WaitForReady() //等待配置中心 获取完成
		if err!=nil{
			errChan<-err
		}
	}()

	r:=gin.New()
	r.Use(CheckForReady())
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Config.JConfig.Data.Mysql})
	})
	go func() {
		err:= r.Run(":8080")
		if err!=nil{
			errChan<-err
		}
	}()
	go func() {
		sig_c:=make(chan os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)
		errChan<-fmt.Errorf("%s",<-sig_c)
	}()
	getErr:=<-errChan
	log.Fatal(getErr)



}