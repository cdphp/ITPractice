import Login from './views/login.vue'
import NotFound from './views/404.vue'
import Index from './views/index.vue'
import Main from './views/layouts/main.vue'
import Invest from './views/invest.vue'
import Master from './views/master.vue'
import User from './views/user/index.vue'
import Article from './views/article/index.vue'

import Register from './views/reg.vue'

let routes = [

    {
        path: '/',
        component: Main,
        name:'',
        children: [
          { path: '/', component: Index, name: '首页', hidden: true },
          { path: '/article', component: Article, name: '文章', hidden: true },
          { path: '/master', component: Master, name: 'Master', hidden: true },
          { path: '/invest', component: Invest, name: '调查', hidden: true },
          { path: '/user', component: User, name: '用户主页', hidden: true },

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
    //{ path: '/main', component: Main },
    /*
    {
        path: '/',
        component: Home,
        name: '导航一',
        iconCls: 'el-icon-message',//图标样式class
        children: [
            { path: '/main', component: Main, name: '主页', hidden: true },
            { path: '/table', component: Table, name: 'Table' },
            { path: '/form', component: Form, name: 'Form' },
            { path: '/user', component: user, name: '列表' },
        ]
    },
    {
        path: '/',
        component: Home,
        name: '导航二',
        iconCls: 'fa fa-id-card-o',
        children: [
            { path: '/page4', component: Page4, name: '页面4' },
            { path: '/page5', component: Page5, name: '页面5' }
        ]
    },
    {
        path: '/',
        component: Home,
        name: '',
        iconCls: 'fa fa-address-card',
        leaf: true,//只有一个节点
        children: [
            { path: '/page6', component: Page6, name: '导航三' }
        ]
    },
    {
        path: '/',
        component: Home,
        name: 'Charts',
        iconCls: 'fa fa-bar-chart',
        children: [
            { path: '/echarts', component: echarts, name: 'echarts' }
        ]
    },
    */
    {
        path: '*',
        hidden: true,
        redirect: { path: '/404' }
    }
];

export default routes;
