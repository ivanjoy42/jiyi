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
      const url = "/api/deck/insert?fronts=" + this.Fronts
      await fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Deck),
      })
      router.back()
    },
  },
}