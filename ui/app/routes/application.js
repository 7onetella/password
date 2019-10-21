import Route from '@ember/routing/route';
import { inject } from '@ember/service';
import ApplicationRouteMixin from 'ember-simple-auth/mixins/application-route-mixin';

export default Route.extend(ApplicationRouteMixin, {
  session: inject('session'),

  setupController (controller, model) {
    console.log("session.isAuthenticated: " + this.get('session.isAuthenticated'));
    this._super(controller, model);
  },

  actions: {
    invalidateSession: function() {
      this.get('session').invalidate();
    }
  }
});