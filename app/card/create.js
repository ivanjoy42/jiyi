export default {
  template: await tpl('card/create'),

  data() {
    return {
      Card: {
        DirId: parseInt(this.$route.query.dirId),
        Front: "",
        Back: "",
        Helper: "",
      },
      DeckId: parseInt(this.$route.query.deckId),
    }
  },

  methods: {
    async postData() {
      const url = "/api/card/insert"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify({ Card: this.Card, DeckId: this.DeckId }),
      })
      router.back()
    },
  },
}