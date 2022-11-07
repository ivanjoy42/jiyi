export default {
  template: await tpl('create'),

  data() {
    return {
      Card: {
        DirId: parseInt(this.$route.query.dirId),
        Front: "",
        Back: "",
        Helper: "",
      },
    }
  },

  methods: {
    async postData() {
      const deckId = this.$route.query.deckId
      const url = "/api/card/insert?deckId=" + deckId
      await fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Card),
      })
      router.back()
    },
  },
}