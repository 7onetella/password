/*eslint no-console: ["error", { allow: ["warn", "error"] }] */
import Route from '@ember/routing/route';

export default Route.extend({  
  model() {
    // var searchtext = "%";
    // console.log("/routes/passwords.js model()");
    // console.log("  searchtext: " + searchtext);

    // reload true results in error when delete is initiated
    // need a way to let force user to wait
    return this.store.findAll("password");
  }
});
