'use strict';

module.exports = function(environment) {
  let ENV = {
    modulePrefix: 'ui',
    environment,
    rootURL: '/ui',
    locationType: 'hash',
    EmberENV: {
      FEATURES: {
        // Here you can enable experimental features on an ember canary build
        // e.g. EMBER_NATIVE_DECORATOR_SUPPORT: true
      },
      EXTEND_PROTOTYPES: {
        // Prevent Ember Data from overriding Date.parse.
        Date: false
      }
    },

    APP: {
      // Here you can pass flags/options to your application instance
      // when it is created
    }
  };

  ENV['ember-simple-auth-token'] = {
    serverTokenEndpoint: '/api/signin',
    serverTokenRefreshEndpoint: '/api/token-refresh',
    tokenPropertyName: 'token', // Key in server response that contains the access token
    headers: {}, // Headers to add to the
    tokenExpirationInvalidateSession: true, // Enables session invalidation on token expiration
    tokenExpireName: 'exp', // Field containing token expiration
    refreshAccessTokens: true,
    refreshLeeway: 60, // refresh 0.1 minutes (10 seconds) before expiration
    refreshTokenPropertyName: 'token', // Key in server response that contains the refresh token
  };

  if (environment === 'development') {
    // ENV.APP.LOG_RESOLVER = true;
    // ENV.APP.LOG_ACTIVE_GENERATION = true;
    // ENV.APP.LOG_TRANSITIONS = true;
    // ENV.APP.LOG_TRANSITIONS_INTERNAL = true;
    // ENV.APP.LOG_VIEW_LOOKUPS = true;

    ENV['ember-simple-auth-token'].serverTokenEndpoint = 'http://localhost:4242/api/signin'
    ENV['ember-simple-auth-token'].serverTokenRefreshEndpoint = 'http://localhost:4242/api/token-refresh'
    ENV.APP.JSONAPIAdaptetHost = 'http://localhost:4242';
  }

  if (environment === 'test') {
    // Testem prefers this...
    ENV.locationType = 'none';

    // keep test console output quieter
    ENV.APP.LOG_ACTIVE_GENERATION = false;
    ENV.APP.LOG_VIEW_LOOKUPS = false;

    ENV.APP.rootElement = '#ember-testing';
    ENV.APP.autoboot = false;
  }

  return ENV;
};
