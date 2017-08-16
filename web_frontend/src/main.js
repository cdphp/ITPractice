import babelpolyfill from 'babel-polyfill'
import Vue from 'vue'
import App from './App'
import ElementUI from 'element-ui'
//import 'element-ui/lib/theme-default/index.css'

import VueRouter from 'vue-router'
import store from './vuex/store'
import Vuex from 'vuex'
//import NProgress from 'nprogress'
//import 'nprogress/nprogress.css'
import routes from './routes'

Vue.use(ElementUI)
Vue.use(VueRouter)
Vue.use(Vuex)



//NProgress.configure({ showSpinner: false });

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  var needLogin = ['/article/add']
  if (needLogin.indexOf(to.path)!=-1) {
    let user = JSON.parse(sessionStorage.getItem('user'));

    if(user) {
      console.log("token:",user)
      if(user.created_at+user.expire < new Date().getTime()/1000) {
        //sessionStorage.removeItem('user');
        //next({ path: '/login' })
      }else {
        next()
      }
    }else {
      next({ path: '/login' })
    }

  }else {
    next()
  }
})

//router.afterEach(transition => {
//NProgress.done();
//});

new Vue({
  //el: '#app',
  //template: '<App/>',
  router,
  store,
  //components: { App }
  render: h => h(App)
}).$mount('#app')
