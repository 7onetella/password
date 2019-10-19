import Component from '@ember/component';
import {inject} from '@ember/service'

export default Component.extend({
  router: inject(),

  actions: {
    edit(password) {
      
      console.log("components.passwords-edit-button id="+password.id);

      this.get('router').transitionTo('/passwords/edit/:id', {
        queryParams: {
          id: password.id
        }
      })
    }
  }
});