import Controller from '@ember/controller';

export default Controller.extend({
  actions: {
    refreshRoute(textsearch) {
      console.log("/controllers/passwords.js refreshRoute()");
      console.log("  textsearch" + textsearch);
      this.sendAction('refreshRoute', textsearch); 
    }
  }
});
