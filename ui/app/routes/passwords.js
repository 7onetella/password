/*eslint no-console: ["error", { allow: ["log", "warn", "error"] }] */
import Route from '@ember/routing/route';

export default Route.extend({

  model() {
    var searchtext = this.get("searchtext");
    console.log("/routes/passwords.js model()");
    console.log("  searchtext: " + searchtext);

    this.store.query('password', {
      filter: {
        title: searchtext
      }
    }).then(function(passwords) {
      return passwords
    });
  }
});
