export default {
  template: await tpl('pract/list'),

  components: { PublicFooter: { template: await tpl('public/footer') } },

  data() {
    return {}
  },
}