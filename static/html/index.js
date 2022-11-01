const App = {
  data() {
    return {}
  }
}

const modify = {
  template: '#modify',
  data() {
    return {
      drawer: false,
      cancel: '取消',
      confirm: '确定',
    }
  },
  beforeRouteLeave() {
    if (this.drawer) {
      this.drawer = false
      return false
    }
  },
}

const routes = [
  { path: '/', component: { template: '#home' } },
  { path: '/list', component: { template: '#list' } },
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
