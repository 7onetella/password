import DS from 'ember-data';
import TokenAuthorizerMixin from 'ember-simple-auth-token/mixins/token-authorizer';
import ENV from '../config/environment';

export default DS.JSONAPIAdapter.extend(TokenAuthorizerMixin, {
  host: ENV.APP.JSONAPIAdaptetHost,
  namespace: 'api'
});
