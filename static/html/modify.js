await tpl('modify.html')

export default {
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