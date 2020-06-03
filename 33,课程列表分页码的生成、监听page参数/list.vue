<template>
  <div>

    <div class="section" >
      <div class="container">
        <div class="tile">
          <div class="tile is-vertical is-parent is-8" id="plist">
            <nav class="pagination is-rounded" role="navigation" aria-label="pagination">
              <a class="pagination-previous" @click="prePage" v-show="this.pageinfo.pageNo!==1">上一页</a>
              <a class="pagination-next" @click="nextPage" v-show="this.pageinfo.pageNo!==this.pageinfo.totalPage">下一页</a>
              <ul class="pagination-list">
                <li v-for="item in this.pageList"   >
                  <a :href="'/course?page='+item.value" v-html="item.value"
                     :class="item.class"></a>
                </li>
              </ul>
            </nav>
            <div class="box" v-for="(item,index) in result">
              <div class="media">
                <div class="media-left">
                  <figure class="image is-64x64">
                    <img src="https://bulma.io/images/placeholders/128x128.png" alt="">
                  </figure>
                </div>
                <div class="media-content lh-20">
                  <p><a href="#"><strong>{{index+1}}、{{item.course_name}}</strong></a> <span class="tag is-danger">推荐</span></p>
                  <p class="is-dark course_intr"  >
                    {{item.course_intr}} </p>

                  <p>380关注 | 68 问题 | 69课堂笔记</p>

                </div>

              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import { getCourseList} from "@/api/CourseAPI";
  export default {
    watchQuery: ['page'],
   data(){
     return {
       pageList: []
     }
   },
   asyncData({error,query}){
     let { page } = query
     if(page==null)
       page = 1
     return getCourseList(page,10).then(res => {
      const { result,pageinfo } = res.data
       if(result ==null || result.length===0){
         error({ statusCode: 400, message: 'no data'})
       }
       return {
        result,
        pageinfo
      }
    }).catch((e) => {
      error({ statusCode: e.code, message: e.message})
    })
  },
    created(){
      this.pageBar();
    },
  methods:{
     pageBar(){
       const {pageNo,pageSize,totalPage} = this.pageinfo
       if(totalPage<=8){
          for(let i=0;i<totalPage;i++){
            this.pageList.push({class:this.getClassName(i+1),value:i+1})
          }
       }else{
         const arr = this.getMiddle(pageNo,totalPage)
         arr.forEach(item => {
           let className = ''
           let value = '&hellip;'
           if(item===0){
             className = 'pagination-ellipsis'
             value ='&hellip;'
             this.pageList.push({class:className,value:value})
           }else {
             this.pageList.push({class:this.getClassName(item),value:item})
           }
         })
       }
     },
    getClassName(index){
      const {pageNo} = this.pageinfo
        return index===pageNo?"pagination-link is-current":"pagination-link"
    },
    getMiddle(current,lastPage){
       let arr=[]
       if(current<=2)
        arr.push(1,2,3,0,lastPage)
       else if(current>=(lastPage-1))
         arr.push(1,0,lastPage-2,lastPage-1,lastPage)
      else
         arr.push(1,0,current-1,current,current+1,0,lastPage)
      return arr
    },
    nextPage(){
      const {pageNo,totalPage} = this.pageinfo
      if(pageNo<totalPage){
        self.location = '/course?page=' + (pageNo+1).toString()
      }
    },
    prePage(){
      const {pageNo,totalPage} = this.pageinfo
      if(pageNo>1){
        self.location = '/course?page=' + (pageNo-1).toString()
      }
    }
  }
}
</script>
<style>
  .course_intr{
    text-indent:1em;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical
  }
  .pagelink .is-current{
     color:white!important;
  }
</style>
