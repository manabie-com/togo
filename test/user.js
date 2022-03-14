process.env.NODE_ENV = 'test';
process.env.MYSQL_DATABASE = 'manabie-test';

const {users} = require('../models').mysql;
let chai = require('chai');
let chaiHttp = require('chai-http');
let server = require('../app');
let should = chai.should();
const jwt = require('jsonwebtoken');
const config = require('../configs');
const jwtSecret = config.getENV('JWT_SECRET') || 'a_secret';
const expiresIn = config.getENV('JWT_EXPIRES_IN') || '1d';
const tokenType = 'Bearer';

chai.use(chaiHttp);

//Our parent block
describe('Users', () => {
  before((done) => {
      //Before test we empty the database in your case
      done();
  });
  /*
  * Test the /GET route
  */
  describe('/POST /user', () => {
    it('it should create new user', (done) => {
      const model = {
          name: "Member",
          email: "member@gmail.com",
          password: "admin123",
          status: "active",
          role: "member"
      }
      const userData = {
        id: 1,
        email: "test@gmail.com",
        role: 'member'
      };
      const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
      chai.request(server)
          .post('/user')
          .set('content-type', 'application/json')
          .set('Authorization', 'Bearer ' + token)
          .send(model)
          .end((err, res) => {
              res.should.have.status(200);
              res.body.data.should.be.a('object');
              res.body.data.should.have.property('id');
              res.body.data.should.have.property('name');
              res.body.data.should.have.property('email');
              res.body.data.should.have.property('role');

              done();
          });
    });
  });

  describe('/GET /users', () => {
    it('it should GET all the users with no Authentication', (done) => {
      chai.request(server)
          .get('/users')
          .end((err, res) => {
              res.should.have.status(401);

              done();
          });
    });
});

  describe('/GET /users', () => {
      it('it should GET all the users', (done) => {
        const userData = {
          id: 1,
          email: "test@gmail.com",
          role: 'member'
        };
        const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
        chai.request(server)
            .get('/users')
            .set('Authorization', 'Bearer ' + token)
            .end((err, res) => {
                res.should.have.status(200);
                res.body.error.should.be.eql(false);
                res.body.data.should.be.a('array');
                done();
            });
      });
  });

  describe('/GET user/:id', () => {
      it('it should GET detail user', (done) => {
        const userData = {
          id: 1,
          email: "test@gmail.com",
          role: 'member'
        };
        const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
        chai.request(server)
            .get('/user/1')
            .set('Authorization', 'Bearer ' + token)
            .end((err, res) => {
                res.should.have.status(200);
                res.body.error.should.be.eql(false);
                res.body.data.should.be.a('object');
                res.body.data.id.should.be.eql(1);
                done();
            });
      });
  });

  describe('/PUT /user/:id', () => {
      it('it should update user', (done) => {
        const model = {
            email: "member_upadte@gmail.com",
        }
        const userData = {
          id: 1,
          email: "test@gmail.com",
          role: 'member'
        };
        const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
        chai.request(server)
            .put('/user/1')
            .set('content-type', 'application/json')
            .set('Authorization', 'Bearer ' + token)
            .send(model)
            .end((err, res) => {
                res.should.have.status(200);
                res.body.data.should.be.a('object');
                res.body.data.should.have.property('id');
                res.body.data.should.have.property('name');
                res.body.data.should.have.property('email');
                res.body.data.should.have.property('role');

                done();
            });
      });
  });

  after((done) => {
    //Before test we empty the database in your case
    users.destroy({
      where: {},
      truncate: true
    });
    done();
  });

});