const supertest = require('supertest');
const config = require('../../../src/config/constants');
const dbQueries = require('../../utils/dbQueries');
const server = require('../../../src/server');
const testData = require('./testData');
const models = require('../../../src/infrastructure/models');
const { testUser } = config;

describe('POST /tasks', () => {
  beforeAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
    await dbQueries.set(testData);
  });

  afterAll(async () => {
    await dbQueries.reset(['tasks', 'users']);
  })

  describe('Create task by user', () => {
    let response;
    let tasks;
    let newTask;

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
        .post('/api/v1/tasks')
        .set('content-type', 'application/json')
        .set('Authorization', `Bearer ${token}`)
        .send({ content: 'new task' });
      
      tasks = await models.tasks.findAll({ where: { user_id: 1 } });
      newTask = await models.tasks.findOne({
        where: { user_id: 1 },
        order: [['id', 'desc']],
      });
    });

    it('status = 200', () => {
      expect(response.status).toEqual(200);
    });

    it('content is created', () => {
      expect(newTask.content).toEqual('new task');
    });

    it('total tasks = 3', () => {
      expect(tasks.length).toEqual(3);
    });
  });
});
