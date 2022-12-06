export default {
  template: await tpl('pract/detail'),

  data() {
    return {
      Pract: {
        PractName: "",
        ModeId: 1,
        DirId: 2,
      },
    }
  },

  created() {
    this.getData()
  },

  methods: {
    async getData() {
      const practId = this.$route.query.practId
      const url = "/api/pract/detail?practId=" + practId
      const json = await (await fetch(url)).json()
      this.Pract = json.Pract
    },
  },
}