export default {
  template: await tpl('dir/list'),

  components: { PublicFooter: { template: await tpl('public/footer') } },

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