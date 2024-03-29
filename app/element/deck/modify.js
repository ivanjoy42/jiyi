export default {
  template: await tpl('deck/modify'),

  data() {
    return {
      Deck: [],
      Fronts: "",
      drawer: false,
    }
  },

  created() {
    this.getData()
  },

  beforeRouteLeave() {
    if (this.drawer) {
      this.drawer = false
      return false
    }
  },

  methods: {
    async postData() {
      const url = "/api/deck/update"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify({ Deck: this.Deck, Fronts: this.Fronts }),
      })
      router.back()
    },

    async getData() {
      const deckId = this.$route.query.deckId
      const url = "/api/deck/modify?deckId=" + deckId
      const json = await (await fetch(url)).json()
      this.Deck = json.Deck
      this.Fronts = json.Fronts
    },

    async delData() {
      const deckId = this.$route.query.deckId
      const url = "/api/deck/delete?deckId=" + deckId
      await fetch(url, {
        method: 'post',
      })
      this.drawer = false
      router.go(-2)
    }
  },
}