import Component from '@ember/component';
import {inject} from '@ember/service'

export default Component.extend({
  router: inject(),

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

      this.sendAction('refreshRoute', searchtext);
    }    
  }
});
