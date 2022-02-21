import { createApp } from 'vue'
import App from './App.vue'
import Wrapper from './Wrapper.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // i just want query string parsing
    { path: '/', component: App }
  ]
})

createApp(Wrapper)
  .use(router)
  .mount('#app')
