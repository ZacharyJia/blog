import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/blog/Index'
import About from '@/blog/About'
import BlogPage from '@/blog/BlogPage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/blog/list/1'
    },
    {
      path: '/blog/list/:page',
      name: 'Index',
      component: Index
    },
    {
      path: '/blog/:id',
      name: 'BlogPage',
      component: BlogPage
    },
    {
      path: '/about',
      name: 'About',
      component: About
    }
  ]
})

