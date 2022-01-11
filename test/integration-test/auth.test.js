const request = require('supertest');
const { MongoMemoryServer } = require("mongodb-memory-server");
const dbConn = require('../../src/utils/db');
const { userModel } = require('../../src/models/model.user');
const app = require('../../src/app');
const mongoose = require('mongoose');

describe("[INTEGRATION TEST]: AUTH TEST.", () => {
  const mongoMock = new MongoMemoryServer();

  beforeAll(async () => {
    await mongoMock.start();
    await dbConn.init(mongoMock.getUri());
  });

  afterEach(async () => {
    console.log('This action running after each.')
    const collections = mongoose.connection.collections;
    for (const item in collections) {
      const collection = collections[item];
      await collection.deleteMany({});
    }
  });

  afterAll(async () => {
    await mongoose.connection.dropDatabase();
    await mongoose.connection.close();
    await mongoMock.stop();
  });

  describe('SIGN-UP API TEST', () => {
    // =================== CASE 01 ===================
    it('Sign up with exists user.', async () => {
      await userModel.create({
        username: 'tiennm',
        password: 'hashedPassword'
      });
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: 'tiennm',
          password: '123456789'
        }).then((res) => {
          expect(res.status).toBe(200); // HTTP API services always return 200 status code. Broswer not throw error.
          expect(res.body.code).toBe(409);
          expect(res.body.message).toBe('username already exists!');
        })
    });

    // =================== CASE 02 ===================
    it('Input username is number.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: 12345678,
          password: 'password'
        }).then((res) => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toEqual(`"username" must be a string`);
        })
    });

    // =================== CASE 03 ===================
    it('Input password is number.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: 'username',
          password: 12345678
        }).then((res) => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toEqual(`"password" must be a string`);
        })
    });

    // =================== CASE 04 ===================
    it('Both username and password are number.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: 1234568,
          password: 12345678
        }).then((res) => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toEqual(`"username" must be a string`);
        })
    });

    // =================== CASE 05 ===================
    it('Input a incorrect field.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: "username",
          password: "password",
          field: 'my_field'
        }).then((res) => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toEqual(`"field" is not allowed`);
        })
    });

    // =================== CASE 06 ===================
    it('Sign up with blank username.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: '',
          password: '12345678'
        }).then((res) => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toEqual(`"username" is not allowed to be empty`);
        })
    });

    // =================== CASE 07 ===================
    it('Sign up with blank password.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: 'username',
          password: ''
        }).then((res) => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.message).toEqual(`"password" is not allowed to be empty`);
        })
    });

    // =================== CASE 07 ===================
    it('Sign up with correct username and password.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send({
          username: 'username',
          password: 'my_password'
        }).then((res) => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(201);
        })
    });
  });

  describe('SIGN-IN API TEST', () => {
    // =================== CASE 01 ===================
    it('Sign-up then sign-in', async () => {
      const user = {
        username: 'tiennm',
        password: 'my_pass_word'
      }

      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send(user)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(201);
        });
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send(user)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(200);
          expect(res.body.success).toBe(true);
          expect(true).toBe(typeof (res.body.data) === 'string');
        });
    });

    // =================== CASE 02 ===================
    it('Sign-in with not exists user.', async () => {
      const user = {
        username: 'tiennm',
        password: 'my_pass_word'
      }

      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send(user)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(404);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Not found user!');
        });
    });

    // =================== CASE 03 ===================
    it('Sign-in with exists user and wrong password.', async () => {
      const user = {
        username: 'tiennm',
        password: 'my_pass_word'
      }

      await request(app.callback())
        .post('/api/public/auth/sign-up')
        .send(user)
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(201);
        });
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send({
          username: 'tiennm',
          password: 'wrong_password'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Invalid username or password!');
        });
    });

    // =================== CASE 04 ===================
    it('Sign-in with an array username.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send({
          username: ['tiennm', 'tiennm01', 'tiennm02'],
          password: 'password'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Invalid input username or password!');
        });
    });

    // =================== CASE 05 ===================
    it('Sign-in with an array password.', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send({
          username: 'tiennm',
          password: ['password01', 'password02']
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Invalid input username or password!');
        });
    });

    // =================== CASE 06 ===================
    it('Sign-in with a number username', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send({
          username: 123456,
          password: 'password'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Invalid input username or password!');
        });
    });

    // =================== CASE 07 ===================
    it('Sign-in with a number password', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send({
          username: 123456,
          password: 'password'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Invalid input username or password!');
        });
    });

    // =================== CASE 08 ===================
    it('Sign-in with blank username', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send({
          username: '',
          password: 'password'
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Invalid input username or password!');
        });
    });

    // =================== CASE 09 ===================
    it('Sign-in with blank passoword', async () => {
      await request(app.callback())
        .post('/api/public/auth/sign-in')
        .send({
          username: 'username',
          password: ''
        })
        .then(res => {
          expect(res.status).toBe(200);
          expect(res.body.code).toBe(400);
          expect(res.body.success).toBe(false);
          expect(res.body.message).toBe('Invalid input username or password!');
        });
    });
  })

})