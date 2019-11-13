const assert = require('assert');
const supertest = require("supertest");
const defaults = require("superagent-defaults");

const { API_ADDR } = process.env;

describe('pre-flight tests', function() {

  describe('OPTIONS /ping', function() {
    it('should be a able to ping with OPTIONS method', function() {
      // note the actual pre-flight OPTIONS doesn't attach custom headers, so we're not using x-api-token here
      // and expect the API to still respect the call
      return defaults(supertest(API_ADDR))
        .options('/ping')
        .expect(200)
        .then(res => { 
          assert(res.headers['access-control-allow-headers'], 'x-api-token')
          assert(res.headers['access-control-allow-methods'], 'GET, POST, PUT, PATCH, DELETE, OPTIONS')
          assert(res.headers['access-control-allow-origin'], '*')
        })
    });

  });

});
