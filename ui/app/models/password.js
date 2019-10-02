import DS from 'ember-data';
const { Model } = DS;

export default Model.extend({
  title: DS.attr('string'),
  url: DS.attr('string'),
  username: DS.attr('string'),
  password: DS.attr('string'),
  notes: DS.attr('string')
});
