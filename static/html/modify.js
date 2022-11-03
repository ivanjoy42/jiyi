export default {
  template: await tpl('modify'),
  data() {
    return {
      drawer: false,
      cancel: '取消',
      confirm: '确定',
      input1: '',
      input2: '',
      input3: '',
    }
  },
  beforeRouteLeave() {
    if (this.drawer) {
      this.drawer = false
      return false
    }
  },
}