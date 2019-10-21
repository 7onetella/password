/*eslint no-console: ["error", { allow: ["log", "warn", "error"] }] */
import Component from '@ember/component';
import {inject as service} from '@ember/service'

export default Component.extend({
  router: service(),
  passwords: service(),
  store: service(),

  actions: {
    edit(password) {
      console.log("/components/passwords.js edit()")
      console.log("  id: "+password.id);

      this.get('router').transitionTo('/passwords/edit/' + password.id)    
    },
    search() {
      console.log("/components/passwords.js search()")
  
      var searchtext = this.get('searchtext');
      console.log("  searchtext: " + searchtext)

      var that = this
      this.store.query('password', {
        filter: {
          title: "%" + searchtext + "%"
        }
      }).then(function(passwords) {
        that.set('password', passwords)
      });
    }    
  }
});
