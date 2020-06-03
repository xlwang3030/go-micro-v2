<template>
  <div>
    列表页
   <div v-for="item in result">
     {{item.course_name}}
   </div>
  </div>
</template>
<script>
  import { getCourseList} from "@/api/CourseAPI";
  export default {
  asyncData(){
    return getCourseList(10).then(res => {
      const { result } = res.data
      return {
        result
      }
    })
  }
}
</script>
