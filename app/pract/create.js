export default {
  template: await tpl('pract/create'),

  data() {
    return {
      Pract: {
        PractName: "",
        ModeId: 1,
        DirId: 2,
      },
      Dir: [],
      Mode: [],
    }
  },

  created() {
    this.getDir()
    this.getMode()
  },

  methods: {
    async getDir() {
      const url = "/api/dir/list"
      const json = await (await fetch(url)).json()
      this.Dir = json.Dir
    },

    async getMode() {
      const url = "/api/mode/list"
      const json = await (await fetch(url)).json()
      this.Mode = json.Mode
    },

    async postData() {
      const url = "/api/pract/insert"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Pract),
      })
      router.back()
    },
  },
}