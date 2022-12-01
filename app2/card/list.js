export default {
  template: await tpl('card/list'),

  data: () => ({
    Card: [],
    Deck: [],
  }),

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const deckId = this.$route.query.deckId
      const dirId = this.$route.query.dirId
      const url = "/api/card/list?dirId=" + dirId + "&deckId=" + deckId
      const json = await (await fetch(url)).json()
      this.Card = json.Card
      this.Deck = json.Deck
    },
  },
}