import VueRouter from "vue-router"
import Vue from "vue"
import CourseList from "~/pages/course/list.vue"
import Index from "~/pages/index.vue"
Vue.use(VueRouter);

const CourseRouterConfig={
  routes:[
    {path:"/",name:"index",
      component:Index},
    {path:"/course",name:"courselist",
      component:CourseList},
  ]
}
export function createRouter() {
  return new VueRouter({
    mode: 'history',
    routes: CourseRouterConfig.routes
  })
}
