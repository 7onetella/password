/*eslint no-console: ["error", { allow: ["warn", "error"] }] */
import Controller from '@ember/controller';
import { inject } from '@ember/service';
import { get, computed } from '@ember/object';
import config from '../config/environment';

export default Controller.extend({
  session: inject('session'),
  userIdle: inject('userIdle'),
  showDebugInfo: false,

  init() {
    this._super(...arguments);
    this.get('userIdle').on('idleChanged', (isIdle) => {
      // isIdle is true if idle. False otherwise.
      // console.log("user idle: " + isIdle)
      if (isIdle === true) {
        this.get('session').invalidate();
      }
    })
  },

  config: computed(function() {
    return JSON.stringify(get(config, 'ember-simple-auth-token'), null, '\t');
  }),

  sessionData: computed('session.session.content.authenticated', function() {
    return JSON.stringify(this.get('session.session.content.authenticated'), null, '\t');
  })
});