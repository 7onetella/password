import Controller from '@ember/controller';
import {inject} from '@ember/service'

export default Controller.extend({
  router: inject(),

  actions: {
    login: function() {
      
      console.log("username: " + $('input[name="username"]')[0].value)
      console.log("password: " + $('input[name="password"]')[0].value)

      $.post("http://localhost:9000/api/signin", {
        username: $('input[name="username"]')[0].value,
        password: $('input[name="password"]')[0].value
      }).then(function() {
        document.location = "/passwords";
      }, function() {
        this.set("loginFailed", true);
      }.bind(this));
    }
  }
});