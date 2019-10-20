import Controller from '@ember/controller';
import {inject} from '@ember/service'

export default Controller.extend({
  router: inject(),

  actions: {
    savePassword(password) {
      console.log("/controllers/passwords/edit.js");

      this.get('router').transitionTo('/passwords')    
    }
  }
});
