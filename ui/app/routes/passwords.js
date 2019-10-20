import Route from '@ember/routing/route';

export default Route.extend({
  model() {
    var searchtext = this.get("searchtext");
    console.log("/routes/passwords.js mode()");
    console.log("  searchtext: " + searchtext);

    return this.store.findAll('password'); 
  },
  actions: {
    refreshRoute(searchtext) {
      console.log("/routes/passwords.js refreshRoute()");
      console.log("  searchtext: " + searchtext);

      this.set('searchtext', searchtext);
      this.refresh();
    }
  }  
});
