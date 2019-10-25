'use strict';

module.exports = function(environment) {
  let ENV = {
    modulePrefix: 'ui',
    environment,
    rootURL: '/ui',
    locationType: 'auto',
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

  if (environment === 'development') {
    // ENV.APP.LOG_RESOLVER = true;
    // ENV.APP.LOG_ACTIVE_GENERATION = true;
    // ENV.APP.LOG_TRANSITIONS = true;
    // ENV.APP.LOG_TRANSITIONS_INTERNAL = true;
    // ENV.APP.LOG_VIEW_LOOKUPS = true;

    ENV['ember-simple-auth-token'] = {
      serverTokenEndpoint: 'https://localhost:4242/api/signin', // Server endpoint to send authenticate request
      tokenPropertyName: 'token', // Key in server response that contains the access token
      headers: {}, // Headers to add to the    
      tokenExpirationInvalidateSession: true, // Enables session invalidation on token expiration
      tokenExpireName: 'exp', // Field containing token expiration      
      refreshAccessTokens: true,
      refreshLeeway: 10, // refresh 0.1 minutes (10 seconds) before expiration
      serverTokenRefreshEndpoint: 'https://localhost:4242/api/token-refresh', // Server endpoint to send refresh request
      refreshTokenPropertyName: 'token', // Key in server response that contains the refresh token
    };     

    ENV.APP.JSONAPIAdaptetHost = 'https://localhost:4242';
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

  if (environment === 'dev') {
    // here you can enable a production-specific feature
    ENV['ember-simple-auth-token'] = {
      serverTokenEndpoint: 'https://dev/api/signin', // Server endpoint to send authenticate request
      tokenPropertyName: 'token', // Key in server response that contains the access token
      headers: {}, // Headers to add to the    
      tokenExpirationInvalidateSession: true, // Enables session invalidation on token expiration
      tokenExpireName: 'exp', // Field containing token expiration      
      refreshAccessTokens: true,
      refreshLeeway: 10, // refresh 0.1 minutes (10 seconds) before expiration
      serverTokenRefreshEndpoint: 'https://dev/api/token-refresh', // Server endpoint to send refresh request
      refreshTokenPropertyName: 'token', // Key in server response that contains the refresh token
    }
  }

  if (environment === 'production') {
    // here you can enable a production-specific feature
    ENV['ember-simple-auth-token'] = {
      serverTokenEndpoint: 'https://keepass/api/signin', // Server endpoint to send authenticate request
      tokenPropertyName: 'token', // Key in server response that contains the access token
      headers: {}, // Headers to add to the    
      tokenExpirationInvalidateSession: true, // Enables session invalidation on token expiration
      tokenExpireName: 'exp', // Field containing token expiration      
      refreshAccessTokens: true,
      refreshLeeway: 10, // refresh 0.1 minutes (10 seconds) before expiration
      serverTokenRefreshEndpoint: 'https://keepass/api/token-refresh', // Server endpoint to send refresh request
      refreshTokenPropertyName: 'token', // Key in server response that contains the refresh token
    }
  }

  return ENV;
};
