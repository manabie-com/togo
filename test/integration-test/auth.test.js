
const mongoose = require("mongoose");
const supertest = require('supertest');
const { expect } = require("chai");
const { MongoMemoryServer } = require("mongodb-memory-server");
const app = require('../../src/app');
const { userModel } = require("../../src/models/model.user");
const port = 9001;
const url = `http://localhost:${port}`;

describe("[INTEGRATION TEST]: AUTH API.", () => {
  const mockDB = new MongoMemoryServer();
  let server = app.listen(port);

  beforeAll(async () => {
    await mockDB.start();
    const mongoUri = mockDB.getUri();
    await mongoose.connect(mongoUri, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });
  });

  afterEach(async () => {
    const collections = mongoose.connection.collections;
    for (const key in collections) {
      const collection = collections[key];
      // await collection.deleteMany({});
    }
  });

  afterAll(async () => {
    await mongoose.connection.dropDatabase();
    await mongoose.connection.close();
    await mockDB.stop();
  });

  describe("[INTEGRATION TEST]: SIGN-UP API.", () => {
    describe("[POST] - /api/public/auth/sign-up", () => {
      // Case 01: Correct data
      it("Sign-up with correct username and password. Then login with username and password.", () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: 'tiennm_001',
            password: 'my_passWord'
          })
          .end((err, res) => {
            if (err) return done(err);
            expect(res.status).equals(200);
            expect(res.body.code).equals(201);
            expect(res?.body?.success).equals(true);
          });
      });

      // Case 02: Input username and password is numebr
      it("Sign-up with in username and password is a number.", () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: 12345678,
            password: 12345678
          })
          .end((err, res) => {
            if (err) return done(err);
            expect(res.status).equals(200);
            expect(res.body.code).equals(400);
            expect(res?.body?.message).includes('must be a string');
          });
      });

      // Case 03: Input a correct field
      it("Sign-up with correct payload.", () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: "username",
            password: "password",
            otherField: 'ACBD'
          })
          .end((err, res) => {
            if (err) return done(err);
            expect(res.status).equals(200);
            expect(res.body.code).equals(400);
            expect(res?.body?.message).includes('is not allowed');
          });
      });

      // Case 4: Input length in-correct
      it("Sign-up with less than string length.", () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: "u",
            password: "p"
          })
          .end((err, res) => {
            if (err) return done(err);
            expect(res.status).equals(200);
            expect(res.body.code).equals(400);
            expect(res?.body?.message).includes('length must be at least');
          });
      });

      // Case 5: Input username already exists.

      it("Sign-up with exists accout.", async () => {
        await userModel.create({
          username: 'exists_account',
          password: '232y78dyS&*ADY*&D^SA'
        });

        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: "exists_account",
            password: "123456"
          })
          .end((err, res) => {
            expect(res.status).equals(200);
            expect(res.body.code).equals(409);
            expect(res?.body?.message).includes('duplicate key error');
          });
      });

      it("Sign-up with blank username", async () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            password: '12345678'
          })
          .end((err, res) => {
            expect(res.status).equals(200);
            expect(res.body.code).equals(400);
            expect(res?.body?.message).includes('"username" is required');
          });
      });

      it("Sign-up with blank password", async () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: 'username'
          })
          .end((err, res) => {
            expect(res.status).equals(200);
            expect(res.body.code).equals(400);
            expect(res?.body?.message).includes('"password" is required');
          });
      });

      it("Sign-up with a number password", async () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: 'username',
            password: 12345678
          })
          .end((err, res) => {
            expect(res.status).equals(200);
            expect(res.body.code).equals(400);
            expect(res?.body?.message).includes('"password" must be a string');
          });
      });

      it("Sign-up with a number username", async () => {
        supertest(server)
          .post(`/api/public/auth/sign-up`)
          .send({
            username: 12345678,
            password: "12345678"
          })
          .end((err, res) => {
            expect(res.status).equals(200);
            expect(res.body.code).equals(400);
            expect(res?.body?.message).includes('"username" must be a string');
          });
      });
    });
  });
});
