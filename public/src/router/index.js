import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/blog/Index'
import About from '@/blog/About'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/about',
      name: 'About',
      component: About
    }
  ]
})

