export default {
  template: await tpl('pract/list'),

  components: { PublicFooter: { template: await tpl('public/footer') } },

  data() {
    return {
      Pract: [],
    }
  },

  created() {
    this.fetchData()
  },

  methods: {
    async fetchData() {
      const url = "/api/pract/list"
      const json = await (await fetch(url)).json()
      this.Pract = json.Pract
    },
  },
}