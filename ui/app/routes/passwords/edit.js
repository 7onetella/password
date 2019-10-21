/*eslint no-console: ["error", { allow: ["log", "warn", "error"] }] */
import Route from '@ember/routing/route';

export default Route.extend({
  model(password) {
    console.log('/routes/passwords/edit.js')
    console.log('  id: '+ password.id)
    return this.store.findRecord('password', password.id);
  }
});
