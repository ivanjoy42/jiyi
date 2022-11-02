const App = {
  data() {
    return {}
  }
}

const routes = [
  { path: '/', component: home },
  { path: '/list', component: list },
  { path: '/modify', component: modify },
]

const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes,
})

const app = Vue.createApp(App)
app.use(ElementPlus, {})
app.use(router)
app.mount("#app")
