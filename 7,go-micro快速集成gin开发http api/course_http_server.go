package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/web"
	"jtthink/src/Course"
	"log"
)

func main()  {
	c:=grpc.NewClient()
	r:=gin.New()
	r.Handle("GET","/test", func(ctx *gin.Context) {
	    c:=Course.NewCourseService("api.jtthink.com.course",c)
	    course_rsp,_:=c.ListForTop(context.Background(),&Course.ListRequest{Size:10})
		ctx.JSON(200,gin.H{"Result":course_rsp.Result})
	})
	service:=web.NewService(
		web.Name("api.jtthink.com.http.course"),
		web.Handler(r),
	 )

	service.Init()
	if err:= service.Run(); err != nil {
		log.Println(err)
	}

}