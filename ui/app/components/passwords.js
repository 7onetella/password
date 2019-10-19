import Component from '@ember/component';

export default Component.extend({
  actions: {
    select(password) {
      console.log('component.passwords.select: '+password.get('title'));
    }
  }
});
