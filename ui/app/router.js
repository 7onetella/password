import EmberRouter from '@ember/routing/router';
import config from './config/environment';

const Router = EmberRouter.extend({
  location: config.locationType,
  rootURL: config.rootURL
});

Router.map(function() {
  this.route('passwords', function() {
    this.route('edit');
  });
  this.route('passwords/new');
  this.route('passwords/edit', { path: '/passwords/edit/:id'});
  this.route('about');
});

export default Router;
