const taskCtrl = require('../handlers/tasksHandler');
const request = require('supertest')
const app = require('../../index');

describe('Post Endpoints', () => {
  test('should create a new post', async (done) => {
    const res = await request(app)
      .post('/tasks')
      .send({
        username: "hungnv2",
        task: "Learn English",
        date: "2022-05-01"
      })
    expect(res.statusCode).toEqual(201);
    done();
  });
  test('should not create a new post', async (done) => {
    const res = await request(app)
      .post('/tasks')
      .send({
        username: "hungnv1",
        task: "Learn English",
        date: "2022-05-01"
      })
    expect(res.statusCode).toEqual(400);
    done();
  });

})