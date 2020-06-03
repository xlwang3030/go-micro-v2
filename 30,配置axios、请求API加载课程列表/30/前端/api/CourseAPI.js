import service from './request'
const api={
  list: '/course'
}
// 获取课程列表（无脑取)
export function getCourseList(size) {
  return service({
    url: api.list,
    method: 'get',
    params: {size}
  })
}
