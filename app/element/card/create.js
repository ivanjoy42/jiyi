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
      Dir: [],
    }
  },

  created() {
    this.getData()
  },

  methods: {
    async getData() {
      const dirId = this.Card.DirId
      const url = "/api/dir/modify?dirId=" + dirId
      const json = await (await fetch(url)).json()
      this.Dir = json.Dir
    },

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