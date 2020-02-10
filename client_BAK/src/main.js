import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'



Vue.config.productionTip = false
Vue.use(VueRouter)

// const Bar = { template: '<div>BAR</div>' }
// const Baz = { template: '<div>BAZ</div>' }

import Manage from './components/admin/manageMain.vue'
import Personal from './components/personal/personalMain.vue'
import Test from './tests/mainTest.vue'

const routes = [
    { path: '/Manage', component: Manage},
    { path: '/Personal', component: Personal},
    { path: '/Test', component: Test }         
]

const router = new VueRouter({
  routes
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
