export default {
  template: await tpl('setting/list'),

  components: { PublicFooter: { template: await tpl('public/footer') } },

  data() {
    return {}
  },
}