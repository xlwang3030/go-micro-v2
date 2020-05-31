package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"

	"jtthink/src/Boot"
	"jtthink/src/Config"

	"os"
	"os/signal"
	"syscall"
	"time"
)


func CheckForReady() gin.HandlerFunc {
	return func(i *gin.Context) {
		if !Boot.ServerIsReady(){
			i.AbortWithStatusJSON(200,gin.H{"result":"server is loading"})
		}else{
			i.Next()
		}
	}
}
func WaitForReady() error  {
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*2)
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


	Boot.BootInit()
	r:=gin.New()
	r.Use(CheckForReady())
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Config.JConfig.Data.Mysql})
	})
	go func() {
		err:= r.Run(":8080")
		if err!=nil{
			Boot.BootErrChan<-err
		}
	}()
	go func() {
		sig_c:=make(chan os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)
		Boot.BootErrChan<-fmt.Errorf("%s",<-sig_c)
	}()
	getErr:=<- Boot.BootErrChan
	logger.Info(getErr)




}