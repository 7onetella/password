/*eslint no-console: ["error", { allow: ["log", "warn", "error"] }] */
import Route from '@ember/routing/route';

export default Route.extend({  
  model() {
    var searchtext = "%";
    console.log("/routes/passwords.js model()");
    console.log("  searchtext: " + searchtext);

    return this.store.findAll("password");
  }
});
