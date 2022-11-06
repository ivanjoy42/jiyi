export default {
  template: await tpl('list'),

  data: () => ({
    Card: [],
    Deck: [],
  }),

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const url = "/api/card/list?dirId=4&deckId=1"
      const json = await (await fetch(url)).json()
      this.Card = json.Card
      this.Deck = json.Deck
    },
  },
}