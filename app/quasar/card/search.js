export default {
  template: await tpl('card/search'),

  data() {
    return {
      Card: [],
      Query: this.$route.query.query,
    }
  },

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const dirId = this.$route.query.dirId
      const url = "/api/card/search?dirId=" + dirId + "&query=" + this.Query
      const json = await (await fetch(url)).json()
      this.Card = json.Card
      this.Query = json.Query
    },

    search() {
      this.fetchData()
      const dirId = this.$route.query.dirId
      router.replace({ path: "search", query: { dirId: dirId, query: this.Query } })
    }
  },
}