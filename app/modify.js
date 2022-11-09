export default {
  template: await tpl('modify'),

  data() {
    return {
      Card: [],
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
    postData() {
      const url = "/api/card/update"
      fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Card),
      })
      router.back()
    },

    async getData() {
      const cardId = this.$route.query.cardId
      const url = "/api/card/modify?cardId=" + cardId
      const json = await (await fetch(url)).json()
      this.Card = json.Card
    },

    async delData() {
      const cardId = this.$route.query.cardId
      const url = "/api/card/delete?cardId=" + cardId
      await fetch(url, {
        method: 'post',
      })
      this.drawer = false
      router.back()
    }
  },
}