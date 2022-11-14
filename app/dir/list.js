export default {
  template: await tpl('dir/list'),

  data: () => ({
    Dir: [],
  }),

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const url = "/api/dir/list"
      const json = await (await fetch(url)).json()
      this.Dir = json.Dir
    },
  },
}