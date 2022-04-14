import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import ArticleList from "../components/ArticleList.vue";
import Detail from "../components/Details.vue"
import Category from '../components/CateList.vue'
import Search from '../components/Search.vue'

Vue.use(VueRouter);


//获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta:{title:'大风起兮'},
    children: [
      {
        path: "/",
        component: ArticleList,
        meta:{title:'大风起兮'},
      },
      {
        path:'detail/:id',
        component: Detail,
        meta:{title:'详情'},
        props:true
      },
      {
        path:'category/:cid',
        component: Category,
        meta:{title:'分类'},
        props:true
      },
      {
        path: 'search/:title',
        component: Search,
        meta: { title: '搜索结果' },
        props: true
      }
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});


router.beforeEach((to,from,next) => {
  //beforeEach是router的钩子函数，在进入路由前执行
  if(to.meta.title){ //有标题时使用标题
     document.title = to.meta.title
  }
  else{
    document.title = '大风起兮'
  }
  next()
})


export default router;
