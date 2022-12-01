export default {
  template: await tpl('card/modify'),

  data() {
    return {
      Card: [],
      Dir: [],
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
    async getData() {
      const cardId = this.$route.query.cardId
      const url = "/api/card/modify?cardId=" + cardId
      const json = await (await fetch(url)).json()
      this.Card = json.Card
      this.Dir = json.Dir
    },

    async postData() {
      const url = "/api/card/update"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Card),
      })
      router.back()
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