export default {
  template: await tpl('user/list'),

  components: { PublicFooter: { template: await tpl('public/footer') } },

  data() {
    return {}
  },
}