export default {
  template: await tpl('dir/modify'),

  data() {
    return {
      Dir: [],
      drawer: false,
    }
  },

  created() {
    this.getData()
  },

  beforeRouteLeave() {
    if (this.drawer) {
      this.drawer = false
      return false
    }
  },

  methods: {
    async postData() {
      const url = "/api/dir/update"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Dir),
      })
      router.back()
    },

    async getData() {
      const dirId = this.$route.query.dirId
      const url = "/api/dir/modify?dirId=" + dirId
      const json = await (await fetch(url)).json()
      this.Dir = json.Dir
    },

    async delData() {
      const dirId = this.$route.query.dirId
      const url = "/api/dir/delete?dirId=" + dirId
      await fetch(url, {
        method: 'post',
      })
      this.drawer = false
      router.go(-2)
    }
  },
}