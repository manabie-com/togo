process.env.NODE_ENV = 'test';
process.env.MYSQL_DATABASE = 'manabie-test';

const {tasks} = require('../models').mysql;
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
describe('Tasks', () => {
  before((done) => {
      //Before test we empty the database in your case
      tasks.destroy({
        where: {},
        truncate: true
      }).then(() => {
        done();
      });
  });
  /*
  * Test the /GET route
  */
  describe('/POST /task', () => {
    it('it should create new task', (done) => {
      const model = {
        name: "Task 1",
        description: "Write API",
        status: "todo",
        estimated_time: 1647047902,
        due_date: 1647047902,
        user_id: 2
      }
      const userData = {
        id: 1,
        email: "test@gmail.com",
        role: 'member'
      };
      const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
      chai.request(server)
          .post('/task')
          .set('content-type', 'application/json')
          .set('Authorization', 'Bearer ' + token)
          .send(model)
          .end((err, res) => {
              res.should.have.status(200);
              res.body.data.should.be.a('object');
              res.body.data.should.have.property('id');
              res.body.data.should.have.property('name');
              res.body.data.should.have.property('status');
              res.body.data.should.have.property('due_date');

              done();
          });
    });
  });

  describe('/GET /tasks', () => {
      it('it should GET all the tasks - with no Authentication', (done) => {
        chai.request(server)
            .get('/tasks')
            .end((err, res) => {
                res.should.have.status(401);
                done();
            });
      });
  });

  describe('/GET /tasks', () => {
      it('it should GET all the tasks', (done) => {
        const userData = {
          id: 1,
          email: "test@gmail.com",
          role: 'member'
        };
        const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
        chai.request(server)
            .get('/tasks')
            .set('Authorization', 'Bearer ' + token)
            .end((err, res) => {
                res.should.have.status(200);
                res.body.error.should.be.eql(false);
                res.body.data.should.be.a('array');
                done();
            });
      });
  });

  describe('/GET task/:id', () => {
      it('it should GET detail task', (done) => {
        const userData = {
          id: 1,
          email: "test@gmail.com",
          role: 'member'
        };
        const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
        chai.request(server)
            .get('/task/1')
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

  describe('/PUT /task/:id', () => {
      it('it should update task', (done) => {
        const model = {
          "status": "in-process"
        }
        const userData = {
          id: 1,
          email: "test@gmail.com",
          role: 'member'
        };
        const token = jwt.sign(userData, jwtSecret, {expiresIn: expiresIn});
        chai.request(server)
            .put('/task/1')
            .set('content-type', 'application/json')
            .set('Authorization', 'Bearer ' + token)
            .send(model)
            .end((err, res) => {
                res.should.have.status(200);
                res.body.data.should.be.a('object');
                res.body.data.should.have.property('id');
                res.body.data.should.have.property('name');
                res.body.data.should.have.property('status');
                res.body.data.should.have.property('due_date');

                done();
            });
      });
  });

});