/*
import Login from './views/login.vue'
import NotFound from './views/404.vue'
import Index from './views/index.vue'
import Main from './views/layouts/main.vue'
import Invest from './views/invest.vue'
import Master from './views/master.vue'
import User from './views/user/index.vue'
import UserEdit from './views/user/edit.vue'
import Article from './views/article/index.vue'
import ArticleAdd from './views/article/add.vue'
import ArticleInfo from './views/article/info.vue'
import Register from './views/reg.vue'
*/
const Register = resolve => require(['./views/reg.vue'], resolve)
const Login = resolve => require(['./views/login.vue'], resolve)
const NotFound = resolve => require(['./views/404.vue'], resolve)
const Index = resolve => require(['./views/index.vue'], resolve)

const Main = resolve => require(['./views/layouts/main.vue'], resolve)
const Master = resolve => require(['./views/master.vue'], resolve)
const User = resolve => require(['./views/user/index.vue'], resolve)
const UserEdit = resolve => require(['./views/user/edit.vue'], resolve)
const Article = resolve => require(['./views/article/index.vue'], resolve)
const ArticleAdd = resolve => require(['./views/article/add.vue'], resolve)
const ArticleInfo = resolve => require(['./views/article/info.vue'], resolve)
let routes = [
    {
        path: '/',
        component: Main,
        name:'',
        children: [
          { path: '/', component: Index, name: '首页', hidden: true },
          { path: '/article', component: Article, name: '文章', hidden: true },
          { path: '/master', component: Master, name: 'Master', hidden: true },
          
          { path: '/user', component: User, name: '用户主页', hidden: true },
          { path: '/user/edit', component: UserEdit, name: '用户编辑', hidden: true },
          { path: '/article/add', component: ArticleAdd, name: '写文章', hidden: true },
          { path: '/article/info', component: ArticleInfo, name: '文章详情', hidden: true },
        ],
    },
    {
        path: '/login',
        component: Login,
        name: '',
        hidden: true
    },
    {
        path: '/reg',
        component: Register,
        name: '',
        hidden: true
    },
    {
        path: '/404',
        component: NotFound,
        name: '',
        hidden: true
    },
    {
        path: '*',
        hidden: true,
        redirect: { path: '/404' }
    }
];

export default routes;
