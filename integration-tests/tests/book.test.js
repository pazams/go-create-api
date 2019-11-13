const assert = require('assert');
const supertest = require("supertest");
const defaults = require("superagent-defaults");

const { API_ADDR, HELPERS_ADDR, API_TOKEN } = process.env;

function apiWithToken(isValid) {
  const request = defaults(supertest(API_ADDR));
  const token = isValid ? API_TOKEN : "foo";
  return request.set('x-api-token', token)
}

function helpersAPI() {
  return supertest(HELPERS_ADDR)
}

describe('integration tests', function() {

  beforeEach(function() {
    return helpersAPI()
      .post('/reset-db')
      .expect(200)

  });

  describe('GET /ping', function() {
    it('should be a able to ping', function() {
      return apiWithToken(true)
        .get('/ping')
        .expect(200)
    });
    it('should NOT be a able to ping with invalid api token', function() {
      return apiWithToken(false)
        .get('/ping')
        .expect(401)
    });
  });

  describe('POST /book', function() {
    it('should be able to insert a book', function() {
      const author = 'pazams';
      const title = 'How To Foo Bar'
      return apiWithToken(true)
        .post('/book')
        .set('Accept', 'application/json')
        .send({
          author,
          title,
        })
        .expect(201)
        .then(res => { 
          assert.deepStrictEqual(res.body.author, author)
          assert.deepStrictEqual(res.body.title, title)
          assert.deepStrictEqual(!!res.body.id, true)
        })
    });
  });

  describe('GET /book', function() {
    it('should be able to get all books', async function() {
      const count = 5
      await createBooks(count)
      return apiWithToken(true)
        .get('/book')
        .set('Accept', 'application/json')
        .expect(200)
        .then(res => { 
          assert.deepStrictEqual(res.body.length, count)
        })
    });
  });

  describe('GET /book/:id', function() {
    it('should be able to get a book by id', async function() {
      await createBooks(10)
      const specialBook = await createBook("cool", "beans")
      await createBooks(10)
      return apiWithToken(true)
        .get(`/book/${specialBook.id}`)
        .expect(200)
        .then(res => { 
          assert.deepStrictEqual(res.body, specialBook)
        })
    });
  });


});


// -------
// HELPERS
// -------
//
async function createBooks(count) {
  for (let i = 0; i < count; i++) {
    await createBook(`author${i}`, `title${i}`);
  }
}

async function createBook(author, title){
  return apiWithToken(true)
    .post('/book')
    .set('Accept', 'application/json')
    .send({
      author,
      title,
    })
    .expect(201)
    .then(res => res.body)
}
