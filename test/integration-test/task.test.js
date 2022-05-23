const request = require('supertest');
const moment = require('moment-timezone');
const { MongoMemoryServer } = require("mongodb-memory-server");
const dbConn = require('../../src/utils/db');
const { userModel } = require('../../src/models/model.user');
const app = require('../../src/app');
const mongoose = require('mongoose');
const { taskModel } = require('../../src/models/model.task');

describe("[INTEGRATION TEST]: TASK TEST.", () => {
  const mongoMock = new MongoMemoryServer();
  let token = null;
  let userId = null;
  let expToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2MWRhNTMwMTVmMjc3ZjNiNWRhYjdiNTUiLCJleHAiOjE2NDE3ODQ4NjQsImlhdCI6MTY0MTY5ODQ2NH0.9RkH2tWay2IO8L7v79IMI1x1mf6hYLsKGZo8Peaus-A';

  beforeAll(async () => {
    await mongoMock.start();
    await dbConn.init(mongoMock.getUri());
    // Create default user for test
    const user = {
      "username": "tiennm",
      "password": "$2b$10$OGo/ZsIDjvirAcbQ9iNnF.RNXk1yU3XslujpQUh0Fjq/q3YTJZxWK", // raw: tiennm
      "limit": 10
    };

    await userModel.create(user).then(res => {
      userId = res._id.toString();
    });
    token = await request(app.callback()).post('/api/public/auth/sign-in').send({
      username: 'tiennm',
      password: 'tiennm'
    }).then(res => {
      return res.body.data;
    });
  });

  afterEach(async () => {
    const collections = mongoose.connection.collections;
    for (const item in collections) {
      if (item !== 'users') { // Not delete user in this test
        const collection = collections[item];
        await collection.deleteMany({});
      }
    }
  });

  afterAll(async () => {
    await mongoose.connection.dropDatabase();
    await mongoose.connection.close();
    await mongoMock.stop();
  });

  describe('[INTEGRATION TEST]: API CREATE TASK', () => {
    // =================== CASE 01 ===================
    it('Create task without token in request', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .send({
          title: 'Title task',
          description: 'Description for this task'
        })
        .then(res => {
          expect(res.status).toBe(401);
          expect(res.body.message).toBe('Request token required!');
        })
    });

    // =================== CASE 02 ===================
    it('Create task with token in request', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task',
          description: 'Description for this task',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(201);
          expect(res.body.success).toBe(true);
        })
    });

    // =================== CASE 03 ===================
    it('Create task with token expired in request', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${expToken}`)
        .send({
          title: 'Title task',
          description: 'Description for this task',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(401);
          expect(res.body.message).toBe('Invalid JWT Token or Token is expired!');
        })
    });

    // =================== CASE 04 ===================
    it('Create task with title is a number.', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 123456,
          description: 'Description for this task',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe('"title" must be a string');
        })
    });

    // =================== CASE 05 ===================
    it('Create task with description is a number.', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task',
          description: 123456,
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe('"description" must be a string');
        })
    });

    // =================== CASE 05 ===================
    it('Create task with status is a string.', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task',
          description: 'Description',
          status: 'str'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe('"status" must be a number');
        })
    });

    // =================== CASE 06 ===================
    it('Create task with un allow fied.', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${token}`)
        .send({
          otherField: 'This is user id',
          title: 'Title task',
          description: 'Description',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe('"otherField" is not allowed');
        })
    });

    // =================== CASE 07 ===================
    it('Create task with all field correct.', async () => {
      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task',
          description: 'Description',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(201);
        })
    });

    // =================== CASE 08 ===================
    it('Create task with over limit daily task.', async () => {
      const listTask = [{
        "userId": userId,
        "title": "Task 001",
        "description": "Description for task 001",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 002",
        "description": "Description for task 002",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 002",
        "description": "Description for task 002",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 003",
        "description": "Description for task 003",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 004",
        "description": "Description for task 004",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 005",
        "description": "Description for task 005",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 006",
        "description": "Description for task 006",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 007",
        "description": "Description for task 007",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 008",
        "description": "Description for task 008",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 009",
        "description": "Description for task 009",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Task 010",
        "description": "Description for task 010",
        "status": 0,
      }];
      await taskModel.insertMany(listTask);

      await request(app.callback())
        .post('/api/me/task')
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task 11',
          description: 'Description 11',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe('Maximum limit daily task!');
        })
    });
  });

  describe('[INTEGRATION TEST]: API UPDATE TASK', () => {
    // =================== CASE 01 ===================
    it('Update not exists task id.', async () => {
      await request(app.callback())
        .put(`/api/me/task?id=61dcf0eee04d62ef4a6ff420`)
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task 11',
          description: 'Description 11',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(404);
          expect(res.body.message).toBe('Not found task id!');
        })
    });
    // =================== CASE 02 ===================
    it(`Update other people's tasks`, async () => {
      let taskId = await taskModel.create({
        userId: '61dcf0eee04d62ef4a6ff420',
        title: 'task title',
        description: 'task description'
      }).then(res => res._id.toString());

      await request(app.callback())
        .put(`/api/me/task?id=${taskId}`)
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task 11',
          description: 'Description 11',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(404);
          expect(res.body.message).toBe('Not found task id!');
        })
    })

    // =================== CASE 03 ===================
    it(`Update owner task.`, async () => {
      let taskId = await taskModel.create({
        userId: userId,
        title: 'Owner task',
        description: 'task description'
      }).then(res => res._id.toString());

      await request(app.callback())
        .put(`/api/me/task?id=${taskId}`)
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task 11',
          description: 'Description 11',
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(200);
          expect(res.body.success).toBe(true);
        })
    });

    // =================== CASE 04 ===================
    it(`Update owner task. With descrription is a number`, async () => {
      let taskId = await taskModel.create({
        userId: userId,
        title: 'Owner task',
        description: 'task description'
      }).then(res => res._id.toString());

      await request(app.callback())
        .put(`/api/me/task?id=${taskId}`)
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 'Title task 11',
          description: 12345,
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe(`"description" must be a string`);
        })
    });

    // =================== CASE 05 ===================
    it(`Update owner task. With title is a number`, async () => {
      let taskId = await taskModel.create({
        userId: userId,
        title: 'Owner task',
        description: 'task description'
      }).then(res => res._id.toString());

      await request(app.callback())
        .put(`/api/me/task?id=${taskId}`)
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: 111,
          description: "12345",
          status: 0
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe(`"title" must be a string`);
        })
    });

    // =================== CASE 06 ===================
    it(`Update owner task. With title is a string`, async () => {
      let taskId = await taskModel.create({
        userId: userId,
        title: 'Owner task',
        description: 'task description'
      }).then(res => res._id.toString());

      await request(app.callback())
        .put(`/api/me/task?id=${taskId}`)
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: "111",
          description: "12345",
          status: 'str'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe(`"status" must be a number`);
        })
    });

    // =================== CASE 07 ===================
    it(`Update owner task. With title an other field not accept.`, async () => {
      let taskId = await taskModel.create({
        userId: userId,
        title: 'Owner task',
        description: 'task description'
      }).then(res => res._id.toString());

      await request(app.callback())
        .put(`/api/me/task?id=${taskId}`)
        .set("Authorization", `Bearer ${token}`)
        .send({
          title: "111",
          description: "12345",
          status: 0,
          otherField: 'blabla'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe(`"otherField" is not allowed`);
        })
    });

    //==================== END GROUP ====================
  });

  // ================== START GROUP LIST ===================
  describe('[INTEGRATION TEST]: API LIST TASK', () => {
    // =================== CASE 01 ===================
    it('GET List user task today.', async () => {
      const listTask = [{
        "userId": '61dcf0eee04d62ef4a6ff420',
        "title": "This other user task.",
        "description": "Description other user task",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Owner Task 001",
        "description": "Description for task 002",
        "status": 0,
      }];
      await taskModel.insertMany(listTask);
      const today = moment().format('YYYY-MM-DD').toString();

      await request(app.callback())
        .get(`/api/me/task?from=${today}&to=${today}`)
        .set("Authorization", `Bearer ${token}`)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(200);
          expect(res.body.data?.length).toBe(1);
        })
    });

    // =================== CASE 02 ===================
    it('GET List deleted task.', async () => {
      const listTask = [{
        "userId": userId,
        "title": "Task deleted",
        "description": "Description other user task",
        "status": -1,
      }, {
        "userId": userId,
        "title": "Owner Task 001",
        "description": "Description for task 002",
        "status": 0,
      }];
      await taskModel.insertMany(listTask);
      const today = moment().format('YYYY-MM-DD').toString();

      await request(app.callback())
        .get(`/api/me/task?from=${today}&to=${today}&status=-1`)
        .set("Authorization", `Bearer ${token}`)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(200);
          expect(res.body.data?.length).toBe(1);
        })
    });

    // =================== CASE 03 ===================
    it('GET List task with page size = 1.', async () => {
      const listTask = [{
        "userId": userId,
        "title": "Task deleted",
        "description": "Description other user task",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Owner Task 001",
        "description": "Description for task 002",
        "status": 0,
      }];
      await taskModel.insertMany(listTask);
      const today = moment().format('YYYY-MM-DD').toString();

      await request(app.callback())
        .get(`/api/me/task?from=${today}&to=${today}&pageSize=1&limit=1`)
        .set("Authorization", `Bearer ${token}`)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(200);
          expect(res.body.data?.length).toBe(1);
        })
    });

    // =================== CASE 03 ===================
    it('GET List task with wrong date time.', async () => {
      const listTask = [{
        "userId": userId,
        "title": "Task deleted",
        "description": "Description other user task",
        "status": 0,
      }, {
        "userId": userId,
        "title": "Owner Task 001",
        "description": "Description for task 002",
        "status": 0,
      }];
      await taskModel.insertMany(listTask);

      await request(app.callback())
        .get(`/api/me/task?from=abcd&to=xyz`)
        .set("Authorization", `Bearer ${token}`)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toBe('Invalid date');
        })
    });
    // ==================== END GROUP LIST ===================
  });

});


