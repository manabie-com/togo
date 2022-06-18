/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

//During the test the env variable is set to test
process.env.NODE_ENV = "test";

//Require the dev-dependencies
const chai = require("chai");
const chaiHttp = require("chai-http");

const server = require("../server");

chai.should();
chai.use(chaiHttp);

describe("Task", () => {
  describe("Unit Test", () => {});
  describe("Integration test", () => {
    it("it should GET list of task of userA (1 task)", (done) => {
      chai
        .request(server)
        .get("/api/tasks")
        .set("Cookie", "access_token=userA:password")
        .end((err, res) => {
          res.should.have.status(200);
          res.body.should.be.a("object");
          res.body.data.should.be.a("array");
          res.body.data.length.should.be.eql(1);

          res.body.data.forEach((task) => {
            task.should.be.a("object");
            task.userId.should.be.a("string");
            task.userId.should.be.eql("A");
          });

          done();
        });
    });
    it("it should GET list of task of userB (0 task)", (done) => {
      chai
        .request(server)
        .get("/api/tasks")
        .set("Cookie", "access_token=userB:password")
        .end((err, res) => {
          res.should.have.status(200);
          res.body.data.should.be.a("array");
          res.body.data.length.should.be.eql(0);
          done();
        });
    });

    it("it should not GET list of task of invalid user", (done) => {
      chai
        .request(server)
        .get("/api/tasks")
        .set("Cookie", "access_token=userD:password")
        .end((err, res) => {
          res.should.have.status(401);
          done();
        });
    });
  });
});
