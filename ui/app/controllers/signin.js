/*eslint no-console: ["error", { allow: ["log", "warn", "error"] }] */
import Controller from '@ember/controller';
import { inject } from '@ember/service'

export default Controller.extend({
  router: inject(),
  session: inject('session'),

  actions: {
    authenticate: function() {
      
      const credentials = this.getProperties('username', 'password');
      console.log(credentials);
      const authenticator = 'authenticator:token'; // or 'authenticator:jwt'

      let promise = this.get('session').authenticate(authenticator, credentials)

      var that = this
      promise.then(function(){
        console.log("  authentication successful. redirecting to listing page");
        that.get('router').transitionTo('/passwords');
      },function(reason) {
        console.log("  reason:" + reason);
        that.set("loginFailed", true);
      });

    }
  }
});