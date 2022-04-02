import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import http from './http'


//rewrite vue 
Vue.prototype.$http = http

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
