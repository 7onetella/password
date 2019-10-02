import DS from 'ember-data';

export default DS.JSONAPIAdapter.extend({
  host: 'http://localhost:9000',
  namespace: 'api'
});
