const request = require('supertest');
const {Todo} = require('../../models/todo');
const {User} = require('../../models/user');
const mongoose = require('mongoose');
const bcrypt = require('bcrypt');
const _ = require('lodash');

let server;

describe('/api/todolist', () => {
  beforeEach(() => { server = require('../../index'); })
  afterEach(async () => { 
    await server.close(); 
    await Todo.remove({});
  });

  describe('POST /', () => {

    let token; 
    let name;
    let user_id = "";

    const exec = async () => {

      return await request(server)
        .post('/api/todolist')
        .set('x-auth-token', token)
        .send({ name, user_id });
    }
    

    beforeEach(() => {
      token = new User().generateAuthToken();      
      name = 'todo1';
    })

    it('should return 401 if client is not logged in', async () => {
      token = ''; 

      const res = await exec();

      expect(res.status).toBe(401);
    });

    it('should return 400 if todo is less than 5 characters', async () => {
      name = '1234'; 
      
      const res = await exec();
      expect(res.status).toBe(400);
    });

    it('should return 400 if todo is more than 50 characters', async () => {
      name = new Array(52).join('a');

      const res = await exec();

      expect(res.status).toBe(400);
    });

    it('should save the todo if it is valid', async () => {
      await exec();

      const todo = await Todo.find({ name: 'todo1' });

      expect(todo).not.toBeNull();
    });

    it('should return 404 if user not exist', async () => {
      user_id = mongoose.Types.ObjectId('4edd40c86762e0fb12000003');

      const res = await exec();
      expect(res.status).toBe(404);
    });
     

  });

});