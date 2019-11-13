const assert = require('assert');
const supertest = require("supertest");
const defaults = require("superagent-defaults");

const { API_ADDR, API_TOKEN } = process.env;

describe('CORS tests', function() {

  it('should be a able to ping with OPTIONS method (without api token)', function() {
    // note the actual pre-flight OPTIONS doesn't attach custom headers, so we're not using x-api-token here
    // and expect the API to still respect the call
    return defaults(supertest(API_ADDR))
      .options('/ping')
      .expect(200)
      .then(res => { 
        assert.deepStrictEqual(res.headers['access-control-allow-headers'], 'x-api-token, Content-Type, Cache-Control')
        assert.deepStrictEqual(res.headers['access-control-allow-methods'], 'GET, POST, PUT, PATCH, DELETE, OPTIONS')
        assert.deepStrictEqual(res.headers['access-control-allow-origin'], '*')
      })
  });

  it('should be a able to ping with GET method (with api token)', function() {
    // note the actual pre-flight OPTIONS doesn't attach custom headers, so we're not using x-api-token here
    // and expect the API to still respect the call
    return defaults(supertest(API_ADDR))
      .set('x-api-token', API_TOKEN)
      .get('/ping')
      .expect(200)
      .then(res => { 
        assert.deepStrictEqual(res.headers['access-control-allow-headers'], 'x-api-token, Content-Type, Cache-Control')
        assert.deepStrictEqual(res.headers['access-control-allow-methods'], 'GET, POST, PUT, PATCH, DELETE, OPTIONS')
        assert.deepStrictEqual(res.headers['access-control-allow-origin'], '*')
      })
  });

});
