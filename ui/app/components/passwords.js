import Component from '@ember/component';

export default Component.extend({
  actions: {
    select(password) {
      console.log('here '+password.get('title'))
    }
  }
});
