/*eslint no-console: ["error", { allow: ["log", "warn", "error"] }] */
import Controller from '@ember/controller';
import {inject} from '@ember/service'

export default Controller.extend({
  router: inject(),

  actions: {
    savePassword(password) {
      console.log("/controllers/passwords/edit.js savePassword()");
      console.log("  tags: "+password.tags);
      console.log("  tags typeof: "+typeof(password.tags));

      this.store.findRecord('password', password.id).then(function(record) {
        // if tags are coming in arrays, don't turn it into array
        if (typeof(password.tags) !== "object") {          
          var tags = password.tags.split(",");
          record.set("tags", tags);
        }
        record.save(); // => PATCH to '/passwords/id'
      });

      this.get('router').transitionTo('/passwords')    
    },
    cancel(password) {
      console.log("/controllers/passwords/edit.js cancel()");
      this.get('router').transitionTo('/passwords');
    },
    deletePassword(password) {
      console.log("/controllers/passwords/edit.js deletePassword()");
      
      if (confirm("Are you sure?")) {
        this.store.findRecord('password', password.id, { backgroundReload: false }).then(function(record) {
          record.destroyRecord(); // => DELETE to '/passwords/id'
        });
    
        this.get('router').transitionTo('/passwords')    
      }
    }  
  },
});
