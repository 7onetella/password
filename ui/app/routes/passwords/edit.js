import Route from '@ember/routing/route';

export default Route.extend({
  model() {
    console.log('routes.passwords/edit')
    return this.store.findAll('password');
  },  
  actions: {
    select() {
      console.log('edit.js id=')
    }
  }
});
