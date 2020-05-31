package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"jtthink/framework/gin_"
	"jtthink/src/Boot"
	"jtthink/src/Config"
	_ "jtthink/src/lib"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main()  {


	Boot.BootInit() //加载各种配置 ，初始化等
	r:=gin.New()
	gin_.BootStrap(r)
	//web改成micro 就是grpc,并直接注册到etcd里面
	service:=web.NewService(
		 web.Name(strings.Join([]string{Config.JConfig.Service.Namespace,Config.JConfig.Service.Name},".")),
		 web.Handler(r),
	 )

	service.Init()
	go func() {
		if err:= service.Run(); err != nil {
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