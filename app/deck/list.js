export default {
  template: await tpl('deck/list'),

  data: () => ({
    Deck: [],
  }),

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const url = "/api/deck/list?dirId=4"
      const json = await (await fetch(url)).json()
      this.Deck = json.Deck
    },
  },
}