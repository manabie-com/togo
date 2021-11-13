const supertest = require('supertest');
const dbQueries = require('../../../utils/dbQueries');
const server = require('../../../../src/server');

describe('POST /register', () => {
  beforeAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
  });

  afterAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
  })

  describe('Register success', () => {
    const defaultMaxTodo = 5;

    let response;
    let result;

    beforeAll(async () => {
      
      response = await supertest
        .agent(server)
        .post('/api/v1/register')
        .set('content-type', 'application/json')
        .send({
          username: 'testUsername',
          password: 'testPassword',
        });
      
      result = JSON.parse(response.text);
    });

    it('status = 200', () => {
      expect(response.status).toEqual(200);
    });

    it('username created', () => {
      expect(result.data.user.username).toEqual('testUsername');
    });

    it('max_todo created', () => {
      expect(result.data.user.max_todo).toEqual(defaultMaxTodo);
    });

    it('returned access token', () => {
      expect(result.data.accessToken).toBeTruthy();
    });
  });

  describe('Register fail', () => {
    let response;
    let result;

    beforeAll(async () => {
      
      response = await supertest
        .agent(server)
        .post('/api/v1/register')
        .set('content-type', 'application/json')
        .send({
          username: 'testUsername',
        });
      
      result = JSON.parse(response.text);
    });

    it('status = 400', () => {
      expect(response.status).toEqual(400);
    });

    it('returned error message', () => {
      expect(result.message).toEqual('Username or password is missing.');
    });
  });
});
