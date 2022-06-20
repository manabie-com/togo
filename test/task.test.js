/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

//During the test the env variable is set to test
process.env.NODE_ENV = "test";

//Require the dev-dependencies
const chai = require("chai");
const chaiHttp = require("chai-http");

const server = require("../server");
const { getTodayString } = require("../utils/index");

chai.should();
chai.use(chaiHttp);

describe("Task", () => {
  describe("Unit Test", () => {
    const TaskMock = require("../repository/Task");

    it("it should add new task", (done) => {
      const numberOfTaskOfUserA = TaskMock.getTaskListByUserId("testA")?.length;

      TaskMock.insertTaskByUserId("testA", { content: "task 2" });

      const numberOfTaskOfUserAAfterAdd =
        TaskMock.getTaskListByUserId("testA")?.length;

      (numberOfTaskOfUserAAfterAdd - numberOfTaskOfUserA).should.be.eql(1);

      done();
    });

    it("it should add new task with right task object", (done) => {
      const taskId = TaskMock.insertTaskByUserId("testA", {
        content: "task 2",
      });

      const task = TaskMock.getTaskById(taskId);

      task.should.be.a("object");
      task.id.should.be.a("string");

      task.content.should.be.a("string");
      task.content.should.be.eql("task 2");

      task.userId.should.be.a("string");
      task.userId.should.be.eql("testA");

      task.createdAt.should.be.a("string");

      done();
    });

    it("it should get number of task of user by created day", (done) => {
      const numberOfTask = 5;
      for (let i = 0; i < numberOfTask; ++i)
        TaskMock.insertTaskByUserId("testB", { content: `task${i}` });

      const result = TaskMock.getNumberOfTaskByUserIdAndDay(
        "testB",
        getTodayString()
      );

      result.should.be.eql(numberOfTask);

      done();
    });
  });

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

    it("it should GET list of task of userB (1 task)", (done) => {
      chai
        .request(server)
        .get("/api/tasks")
        .set("Cookie", "access_token=userB:password")
        .end((err, res) => {
          res.should.have.status(200);
          res.body.data.should.be.a("array");
          res.body.data.length.should.be.eql(1);
          done();
        });
    });

    it("it should GET list of task of userC (0 task)", (done) => {
      chai
        .request(server)
        .get("/api/tasks")
        .set("Cookie", "access_token=userC:password")
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

    it("it should add new task for userA", (done) => {
      chai
        .request(server)
        .post("/api/task")
        .send({ content: "task 2" })
        .set("Cookie", "access_token=userA:password")
        .end((err, res) => {
          res.should.have.status(200);
          res.body.data.should.be.a("object");
          res.body.data.taskId.should.be.a("string");
          done();
        });
    });

    it("it should not add task for userA (lack content)", (done) => {
      chai
        .request(server)
        .post("/api/task")
        .set("Cookie", "access_token=userA:password")
        .end((err, res) => {
          res.should.have.status(200);
          res.body.error_code.should.be.a("number");
          res.body.error_code.should.be.eql(10000);
          res.body.error_message.should.be.a("string");
          res.body.error_message.should.be.eql("Invalid task");
          done();
        });
    });

    it("it should update list of task of userA (2 task)", (done) => {
      chai
        .request(server)
        .get("/api/tasks")
        .set("Cookie", "access_token=userA:password")
        .end((err, res) => {
          res.should.have.status(200);
          res.body.should.be.a("object");
          res.body.data.should.be.a("array");
          res.body.data.length.should.be.eql(2);

          res.body.data.forEach((task) => {
            task.should.be.a("object");
            task.userId.should.be.a("string");
            task.userId.should.be.eql("A");
          });

          done();
        });
    });

    it("it should not add task for userC (reach the limit)", (done) => {
      chai
        .request(server)
        .post("/api/task")
        .send({ content: "task 3" })
        .set("Cookie", "access_token=userC:password")
        .end((err, res) => {
          res.should.have.status(200);
          res.body.error_code.should.be.a("number");
          res.body.error_code.should.be.eql(10001);
          res.body.error_message.should.be.a("string");
          res.body.error_message.should.be.eql("Reach the limit rate");
          done();
        });
    });

    it("it should not add task for invalid user", (done) => {
      chai
        .request(server)
        .post("/api/task")
        .send({ content: "task 3" })
        .set("Cookie", "access_token=userD:password")
        .end((err, res) => {
          res.should.have.status(401);
          done();
        });
    });
  });
});
