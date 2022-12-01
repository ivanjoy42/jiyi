export default {
  template: await tpl('pract/create'),

  data() {
    return {
      Pract: {
        PractName: "",
        ModeId: 1,
        DirId: 2,
      },
      Mode: [],
      Dir: [],
      ModeName: "",
      DirName: "",
      modeDrawer: false,
      dirDrawer: false,
    }
  },

  beforeRouteLeave() {
    if (this.modeDrawer) {
      this.modeDrawer = false
      return false
    }

    if (this.dirDrawer) {
      this.dirDrawer = false
      return false
    }
  },

  created() {
    this.getDir()
    this.getMode()
  },

  methods: {
    setMode(ModeName) {
      this.ModeName = ModeName
      this.modeDrawer = false
      this.Pract.PractName = this.ModeName + this.DirName
    },

    setDir(DirName) {
      this.DirName = DirName
      this.dirDrawer = false
      this.Pract.PractName = this.ModeName + this.DirName
    },

    async getDir() {
      const url = "/api/dir/list"
      const json = await (await fetch(url)).json()
      this.Dir = json.Dir
    },

    async getMode() {
      const url = "/api/mode/list"
      const json = await (await fetch(url)).json()
      this.Mode = json.Mode
    },

    async postData() {
      const url = "/api/pract/insert"
      await fetch(url, {
        method: 'post',
        body: JSON.stringify(this.Pract),
      })
      router.back()
    },
  },
}