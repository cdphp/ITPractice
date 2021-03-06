
const Index = resolve => require(['./views/index.vue'], resolve)
const Login = resolve => require(['./views/common/login.vue'], resolve)
const Register = resolve => require(['./views/register/index.vue'], resolve)
const Validate = resolve => require(['./views/common/validate.vue'], resolve)
const Forget = resolve => require(['./views/common/forget.vue'], resolve)
const NotFound = resolve => require(['./views/common/404.vue'], resolve)
const Oauth = resolve => require(['./views/common/oauth.vue'], resolve)

const Main = resolve => require(['./views/layouts/main.vue'], resolve)
const Master = resolve => require(['./views/master.vue'], resolve)
const User = resolve => require(['./views/user/index.vue'], resolve)
const UserEdit = resolve => require(['./views/user/edit.vue'], resolve)
const Article = resolve => require(['./views/article/index.vue'], resolve)
const ArticleAdd = resolve => require(['./views/article/add.vue'], resolve)
const ArticleInfo = resolve => require(['./views/article/info.vue'], resolve)
const ArticleEdit = resolve => require(['./views/article/edit.vue'], resolve)
const Company = resolve => require(['./views/company/index.vue'], resolve)
const CompanyAdd = resolve => require(['./views/company/add.vue'], resolve)
const Live = resolve => require(['./views/live/info.vue'], resolve)
const Question = resolve => require(['./views/question/index.vue'], resolve)
const QuestionAdd = resolve => require(['./views/question/add.vue'], resolve)
const QuestionInfo = resolve => require(['./views/question/info.vue'], resolve)
const QuestionEdit = resolve => require(['./views/question/edit.vue'], resolve)
let routes = [
    {
        path: '/',
        component: Main,
        name:'',
        children: [
          { path: '/', component: Index, name: '首页', hidden: true },
          { path: '/article', component: Article, name: '文章', hidden: true },
          { path: '/company', component: Company, name: '公司', hidden: true },
          { path: '/master', component: Master, name: '传承', hidden: true },
          { path: '/live', component: Live, name: '直播', hidden: true },
          { path: '/question', component: Question, name: '问道', hidden: true },
          { path: '/oauth', component: Oauth, name: 'Oauth', hidden: true },

          { path: '/user', component: User, name: '用户主页', hidden: true },
          { path: '/user/edit', component: UserEdit, name: '用户编辑', hidden: true },
          { path: '/company/add', component: CompanyAdd, name: '添加公司', hidden: true },
          { path: '/article/add', component: ArticleAdd, name: '写文章', hidden: true },
          { path: '/article/info', component: ArticleInfo, name: '文章详情', hidden: true },
          { path: '/article/edit', component: ArticleEdit, name: '文章编辑', hidden: true },
          { path: '/reg', component: Register, name: '注册', hidden: true },
          { path: '/validate', component: Validate, name: '等待验证', hidden: true },

          { path: '/login', component: Login, name: '登录', hidden: true },
          { path: '/forget', component: Forget, name: '忘记密码', hidden: true },
          { path: '/question/add', component: QuestionAdd, name: '发问', hidden: true },
          { path: '/question/info', component: QuestionInfo, name: '问题详情', hidden: true },
          { path: '/question/edit', component: QuestionEdit, name: '问题编辑', hidden: true },
        ],
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
