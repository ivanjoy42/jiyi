export default {
  template: await tpl('list'),

  data: () => ({
    cardList: null,
  }),

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const url = "list.json"
      this.cardList = await (await fetch(url)).json()
    },
  },
}