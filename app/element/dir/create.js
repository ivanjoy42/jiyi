export default {
  template: await tpl('dir/create'),

  data() {
    return {
      Dir: {
        DirName: "",
        KindId: 1,
      },
    }
  },

  methods: {
    async postData() {
      const url = "/api/dir/insert"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Dir),
      })
      router.back()
    },
  },
}