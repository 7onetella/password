/*eslint no-console: ["error", { allow: ["warn", "error"] }] */
import Controller from '@ember/controller';
import {inject} from '@ember/service'

export default Controller.extend({
  router: inject(),

  actions: {
    savePassword(password) {
      // console.log("/controllers/passwords/new.js");

      let record = this.store.createRecord('password', password);
      
      record.save(); // => POST to '/passwords'

      this.get('router').transitionTo('/passwords')    
    }
  }
});
