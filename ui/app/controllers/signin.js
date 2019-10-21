/*eslint no-console: ["error", { allow: ["log", "warn", "error"] }] */
import Controller from '@ember/controller';
import { inject } from '@ember/service'

export default Controller.extend({
  router: inject(),
  session: inject('session'),

  actions: {
    authenticate: function() {
      
      console.log("username: " + this.get("username"))
      console.log("password: " + this.get("password"))

      $.post("http://localhost:9000/api/signin", {
        username: this.get("username"),
        password: this.get("password"),
      }).then(function() {
        // console.log("signin successful")
        // document.location = "/passwords";
      }, function(data) {
        console.log(data);
        if (data.status == 200) {
          document.location = "/passwords";
        } else {
          this.set("loginFailed", true);
        }
      }.bind(this));

    }
  }
});