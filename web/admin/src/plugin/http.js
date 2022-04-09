import Vue from 'vue'
import axios from 'axios'

let Url = 'http://localhost:3000/api/v1/'

axios.defaults.baseURL = Url


axios.interceptors.request.use(config => { //请求拦截器 详见axios文档 设置携带token
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios

export { Url }