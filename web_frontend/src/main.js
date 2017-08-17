
import Vue from 'vue'
import { Form,FormItem,Button,Checkbox,Input,Message,MessageBox } from 'element-ui'
import App from './App.vue'
import VueRouter from 'vue-router';

Vue.use(Form)
Vue.use(FormItem)
Vue.use(Button)
Vue.use(Checkbox)
Vue.use(Input)

Vue.prototype.$confirm = MessageBox.confirm

Vue.prototype.$message = Message
Vue.use(VueRouter)

import routes from './routes'
const router = new VueRouter({
  routes
})



new Vue({
  router,
  el: '#app',
  render: h => h(App)
}).$mount('#app')
