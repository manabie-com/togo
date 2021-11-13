const supertest = require('supertest');
const config = require('../../../../src/config/constants');
const dbQueries = require('../../../utils/dbQueries');
const server = require('../../../../src/server');
const masterdata = require('../../../utils/masterdata');
const { testUser } = config;

describe('POST /login', () => {
  beforeAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
    await dbQueries.set(masterdata);
  });

  afterAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
  })

  describe('Login success', () => {
    const currentUser = masterdata.users[0];

    let response;
    let result;

    beforeAll(async () => {
      response = await supertest
        .agent(server)
        .post('/api/v1/login')
        .set('content-type', 'application/json')
        .send({
          username: testUser.username,
          password: testUser.password,
        });
      
      result = JSON.parse(response.text);
    });

    it('status = 200', () => {
      expect(response.status).toEqual(200);
    });

    it('username invalid', () => {
      expect(result.data.user.username).toEqual(currentUser.username);
    });

    it('max_todo invalid', () => {
      expect(result.data.user.max_todo).toEqual(currentUser.max_todo);
    });

    it('returned access token', () => {
      expect(result.data.accessToken).toBeTruthy();
    });
  });

  describe('Login fail', () => {
    let response;
    let result;

    beforeAll(async () => {
      
      response = await supertest
        .agent(server)
        .post('/api/v1/login')
        .set('content-type', 'application/json')
        .send({
          username: 'testUsername',
          password: 'testPass',
        });
      
      result = JSON.parse(response.text);
    });

    it('status = 400', () => {
      expect(response.status).toEqual(400);
    });

    it('returned error message', () => {
      expect(result.message).toEqual('Username or password is incorrect.');
    });
  });
});
