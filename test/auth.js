process.env.NODE_ENV = 'test';
process.env.MYSQL_DATABASE = 'manabie-test';

const {users} = require('../models').mysql;
let chai = require('chai');
let chaiHttp = require('chai-http');
let server = require('../app');
let should = chai.should();
const bcrypt = require('bcrypt');
const config = require('../configs');
const saltRounds = config.getENV('SALT_ROUNDS');

describe('Auth', () => {
  before((done) => {
      //Before test we empty the database in your case
      done();
  });

  describe('/POST /login', () => {
    it('it should user login - password is invalid', (done) => {
      const model = {
        email: "test@gmail.com",
        password: "123admin",
      }
      chai.request(server)
          .post('/login')
          .set('content-type', 'application/json')
          .send(model)
          .end((err, res) => {
              res.should.have.status(401);

              done();
          });
    });
  });

  describe('/POST /login', () => {
    it('it should user login - user not exist', (done) => {
      const model = {
        email: "test123@gmail.com",
        password: "admin123",
      }
      chai.request(server)
          .post('/login')
          .set('content-type', 'application/json')
          .send(model)
          .end((err, res) => {
              res.should.have.status(401);

              done();
          });
    });
  });

  describe('/POST /login', () => {
    it('it should user login - success', (done) => {
      const model = {
        email: "test@gmail.com",
        password: "admin123",
      }
      chai.request(server)
          .post('/login')
          .set('content-type', 'application/json')
          .send(model)
          .end((err, res) => {
              res.should.have.status(200);
              res.body.data.should.be.a('object');
              res.body.data.should.have.property('token');

              done();
          });
    });
  });
});