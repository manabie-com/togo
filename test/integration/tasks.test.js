process.env.NODE_ENV = "test";

const mongoose = require("mongoose");
const chaiHttp = require("chai-http");
const chai = require("chai");
const Tasks = require("../../src/models/tasks.model");
const Users = require("../../src/models/users.model");
const server = require("../../server");
const should = chai.should();

let testingUser = {
  username: "admin",
  password: "admin@123"
};

let accessToken, userId, taskId;

chai.use(chaiHttp);
describe("Test Task apis", () => {
  before(done => { //register user for testing
    Users.deleteMany({}, (err) => {
      done();
    });
  });
  before(done => { //register user for testing
    chai
      .request(server)
      .post("/api/users/register")
      .send(testingUser)
      .end((err, res) => {
        res.should.have.status(200);
        done();
      });
  });
  before(done => { //login and get token
    chai
      .request(server)
      .post("/api/login")
      .send(testingUser)
      .end((err, res) => {
        accessToken = res.body.token;
        userId = res.body.user._id;
        res.should.have.status(200);
        done();
      });
  });
  after(done => {
    Users.deleteMany({}, (err) => {
      done();
    });
  });
  after(done => {
    Tasks.deleteMany({}, (err) => {
      done();
    });
  });
  describe("/GET tasks", () => {
    it("should GET all the tasks", (done) => {
      chai.request(server)
        .get("/api/tasks")
        .set("Authorization", accessToken)
        .end((err, res) => {
          res.should.have.status(200);
          res.body.should.be.a("array");
          res.body.length.should.be.eql(0);
          done();
        });
    });
  });

  describe("/POST tasks", () => {
    beforeEach(done => { //set limit a day for user
      chai
        .request(server)
        .put("/api/users/" + userId)
        .set("Authorization", accessToken)
        .send({ limit: 1 })
        .end((err, res) => {
          res.should.have.status(200);
          done();
        });
    });
    it("should POST a task based on limit of user", (done) => {
      chai.request(server)
        .post("/api/tasks")
        .set("Authorization", accessToken)
        .send({ name: "To go", username: "admin" })
        .end((err, res) => {
          res.should.have.status(200);
          res.body.should.be.a("object");
          res.body.should.have.property("name").eql("To go");
          res.body.should.have.property("user").eql(userId);
          taskId = res.body._id;
          done();
        });
    });
    it("should throw a error when reached limit of tasks per day", (done) => {
      chai.request(server)
        .post("/api/tasks")
        .set("Authorization", accessToken)
        .send({ name: "To go 2", username: "admin" })
        .end((err, res) => {
          res.should.have.status(500);
          done();
        });
    });
  });

  describe("/PUT tasks", () => {
    it("should PUT a task by id", (done) => {
      chai.request(server)
        .put("/api/tasks/" + taskId)
        .set("Authorization", accessToken)
        .send({ name: "To go 123" })
        .end((err, res) => {
          res.should.have.status(200);
          res.body.should.be.a("object");
          res.body.should.have.property("name").eql("To go 123");
          done();
        });
    });
  });

  describe("/DELETE tasks", () => {
    it("should DELETE a task by id", (done) => {
      chai.request(server)
        .delete("/api/tasks/" + taskId)
        .set("Authorization", accessToken)
        .end((err, res) => {
          res.should.have.status(200);
          res.body.should.be.a("object");
          res.body.should.have.property("_id").eql(taskId);
          done();
        });
    });
  });
});
