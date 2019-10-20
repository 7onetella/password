import Controller from '@ember/controller';
import {inject} from '@ember/service'

export default Controller.extend({
  router: inject(),

  actions: {
    savePassword(password) {
      console.log("/controllers/passwords/edit.js");

      this.store.findRecord('password', password.id).then(function(record) {
        record.save(); // => PATCH to '/passwords/id'
      });

      this.get('router').transitionTo('/passwords')    
    },
    deletePassword(password) {
      console.log("/controllers/passwords/edit.js");
  
      this.store.findRecord('password', password.id, { backgroundReload: false }).then(function(record) {
        record.destroyRecord(); // => DELETE to '/passwords/id'
      });
  
      this.get('router').transitionTo('/passwords')    
    }  
  },
});
