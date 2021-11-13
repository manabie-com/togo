const supertest = require('supertest');
const config = require('../../../src/config/constants');
const dbQueries = require('../../utils/dbQueries');
const server = require('../../../src/server');
const testData = require('./testData');
const { testUser } = config;

describe('GET /tasks', () => {
  beforeAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
    await dbQueries.set(testData);
  });

  afterAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
  })

  describe('Get list tasks by user', () => {
    let response;
    let result;

    beforeAll(async () => {
      const resLogin = await supertest
        .agent(server)
        .post('/api/v1/login')
        .set('content-type', 'application/json')
        .send({
          username: testUser.username,
          password: testUser.password,
        });
      const token = JSON.parse(resLogin.text).data.accessToken;
      
      response = await supertest
        .agent(server)
        .get('/api/v1/tasks')
        .set('content-type', 'application/json')
        .set('Authorization', `Bearer ${token}`);
      
      result = JSON.parse(response.text);
    });

    it('status = 200', () => {
      expect(response.status).toEqual(200);
    });

    it('total tasks = 2', () => {
      expect(result.data.length).toEqual(2);
    });
  });
});
