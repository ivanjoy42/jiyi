export default {
  template: await tpl('deck/create'),

  data() {
    return {
      Deck: {
        DirId: parseInt(this.$route.query.dirId),
        DeckName: "",
      },
      Fronts: "",
    }
  },

  methods: {
    async postData() {
      const url = "/api/deck/insert"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify({ Deck: this.Deck, Fronts: this.Fronts }),
      })
      router.back()
    },
  },
}