export default {
  template: await tpl('deck/list'),

  data: () => ({
    Deck: [],
    Dir: [],
  }),

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const dirId = this.$route.query.dirId
      const url = "/api/deck/list?dirId=" + dirId
      const json = await (await fetch(url)).json()
      this.Deck = json.Deck
      this.Dir = json.Dir
    },
  },
}