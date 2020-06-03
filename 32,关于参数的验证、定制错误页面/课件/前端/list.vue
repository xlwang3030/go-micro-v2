<template>
  <div>

    <div class="section" >
      <div class="container">
        <div class="tile">
          <div class="tile is-vertical is-parent is-8" id="plist">
            <nav class="pagination is-rounded" role="navigation" aria-label="pagination">
              <a class="pagination-previous">Previous</a>
              <a class="pagination-next">Next page</a>
              <ul class="pagination-list">
                <li><a class="pagination-link" aria-label="Goto page 1">1</a></li>
                <li><span class="pagination-ellipsis">&hellip;</span></li>
                <li><a class="pagination-link" aria-label="Goto page 45">45</a></li>
                <li><a class="pagination-link is-current" aria-label="Page 46" aria-current="page">46</a></li>
                <li><a class="pagination-link" aria-label="Goto page 47">47</a></li>
                <li><span class="pagination-ellipsis">&hellip;</span></li>
                <li><a class="pagination-link" aria-label="Goto page 86">86</a></li>
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

   asyncData({error}){
    return getCourseList(-1,10).then(res => {
      const { result } = res.data
      return {
        result
      }
    }).catch((e) => {
      error({ statusCode: e.code, message: e.message})
    })
  },
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
</style>
