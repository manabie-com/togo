`use strict`

  //During the test the env variable is set to test
process.env.NODE_ENV = 'test';

//Require the dev-dependencies
let chai = require('chai');
let chaiHttp = require('chai-http');
let server = require('../server');
let should = chai.should();

chai.use(chaiHttp);

describe('/POST task', () => {
  it('it should POST a task', (done) => {
      let task = {
        userID: "001",
        data: "Morning exercise"
      };
      chai.request(server)
          .post('/v1/create')
          .send(task)
          .end((err, res) => {
              res.should.have.status(200);
              res.body.should.be.a('object');
              res.body.should.have.property('message').eql('Task successfully added!');
              res.body.pet.should.have.property('id');
              res.body.pet.should.have.property('userID').eql(task.userID);
              res.body.pet.should.have.property('data').eql(task.data);
              done();
          });
  });
  it('it should not POST a task without task data', (done) => {
      let task = {
        userID: "001"
      };
      chai.request(server)
          .post('/v1/create')
          .send(task)
          .end((err, res) => {
              res.should.have.status(200);
              res.body.should.be.a('object');
              res.body.should.have.property('message').eql("Invalid task data");
              done();
          });
  });
});