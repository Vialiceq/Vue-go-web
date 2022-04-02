import Vue from 'vue'
import VueRouter from 'vue-router'
//import Home from '../views/Home.vue'
import Main from '../views/Main.vue'
//import About from '../views/About.vue'
import Categories from '../views/Categories.vue'
import CategoriesList from '../views/CategoriesList.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main,
    children: [
      { path: '/categories/create', component: Categories },
      { path: '/categories/list', component: CategoriesList }
  
  ]
  },

  
]

const router = new VueRouter({
  routes
})

export default router
