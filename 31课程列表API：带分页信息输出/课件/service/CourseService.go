package service

import (
	"context"
	. "jtthink/src/Course"
	"jtthink/src/Mapper"
)
func NewCourseModel(id int32,name string) *CourseModel  {
	return &CourseModel{CourseId:id,CourseName:name}
}
type CourseServiceImpl struct{}
func(this *CourseServiceImpl) ListForTop(ctx context.Context, req *ListRequest, rsp *ListResponse) error{
	 course:=make([]*CourseModel,0)
	 totalCount:=0
	 db:=Mapper.GetCourseListBySql(req.Size,(req.Page-1)*req.Size)
	 err:=db.Find(&course).Error
	if err!=nil{
		return err
	}
    Mapper.GetCourseListCount(db,&totalCount)
	rsp.Result=course
	rsp.Pageinfo=BuildPageInfo(req.Page,req.Size,int32(totalCount))
	return nil
}
func(this *CourseServiceImpl) GetDetail(ctx context.Context,req *DetailRequest, rsp *DetailResponse) error{
	//只取课程详细
	if req.FetchType==0 || req.FetchType==1 || req.FetchType==3{
		if err:=Mapper.GetCourseDetail(int(req.CourseId)).Find(rsp.Course).Error;err!=nil{
			return err
		}
	}
	//只取计数表详细
	if  req.FetchType==2 || req.FetchType==3{
		if err:=Mapper.GetCourseCounts(int(req.CourseId)).Find(&rsp.Counts).Error;err!=nil{
			return err
		}
	}

	return nil

}

func NewCourseServiceImpl() *CourseServiceImpl  {
	return &CourseServiceImpl{}
}
func BuildPageInfo(page int32,size int32,count int32 ) *PageInfo {
	pageInfo:=&PageInfo{PageNo:page,PageSize:size,TotalCount:count}
	var pageNum int32=1
	if count>size{
		pageNum=count/size
		if count%size>0{
			pageNum++
		}
	}
	pageInfo.TotalPage=pageNum
	return pageInfo
}
