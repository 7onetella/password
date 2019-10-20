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