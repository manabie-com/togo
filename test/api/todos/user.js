
const app = require('../../../main');
const chai = require("chai");
const chaiHttp = require("chai-http");
//Assertion Style
chai.should();

chai.use(chaiHttp);

describe('Todos API', () => {
    /**
     * Test the GET route
     */
    describe("GET /api/v1/todo", () => {
        it("It should GET all the todo", (done) => {
            chai.request(app)
                .get("/api/v1/todo")
                .end((err, response) => {
                    response.should.have.status(200);
                    response.body.should.be.a("object");
                    done();
                });
        });
    });

    /**
     * Test the GET (by id) route
     */
    describe("GET /api/v1/todo/:id", () => {
        it("It should GET a todo by ID", (done) => {
            const todoId = "619a43ff2ecf969a9784945c";
            chai.request(app)
                .get("/api/v1/todo/" + todoId)
                .end((err, response) => {
                    response.should.have.status(200);
                    response.body.should.be.a("object");
                    response.body.should.have.property("success").eq(true);
                    done();
                });
        });

        it("It should NOT GET a todo by ID", (done) => {
            const todoId = 123;
            chai.request(app)
                .get("/api/v1/todo/" + todoId)
                .end((err, response) => {
                    response.should.have.status(404);
                    response.body.should.have.property("message").eq("Invalid Id");
                    done();
                });
        });

    });



    /**
     * Test the POST route
     */


    /**
    * Test the PATCH route
    */
    describe("PATCH /api/v1/todo/:id", () => {
        it("It should PATCH an existing todo", (done) => {
            const todoId = "619a43ff2ecf969a9784945c";
            const todo = {
                content: "this is content 1"
            };
            chai.request(app)
                .patch("/api/v1/todo/" + todoId)
                .send(todo)
                .end((err, response) => {
                    response.should.have.status(200);
                    response.body.should.be.a('object');
                    response.body.should.have.property("sucess").eq(true);
                    done();
                });
        });

        it("It should NOT PATCH an existing task with a name with less than 3 characters", (done) => {
            const todoId = "619a43ff2ecf969a9784945c";
            const todo = {
                content: " "
            };
            chai.request(app)
                .patch("/api/v1/todo/" + todoId)
                .send(todo)
                .end((err, response) => {
                    response.should.have.status(404);
                    response.body.should.have.property("message").eq("Must be at least 3 characters");
                    done();
                });
        });
    });



    /**
    * Test the DELETE (by id) route
    */

    describe("DELETE /api/v1/todo/:id", () => {
        it("It should DELETE an existing todo", (done) => {
            const todoId = "619a43ff2ecf969a9784945c";
            chai.request(app)
                .delete("/api/v1/todo/" + todoId)
                .end((err, response) => {
                    response.should.have.status(200);
                    response.body.should.have.property("success").eq(true);
                    done();
                });
        });

        it("It should NOT DELETE a todo that is not in the database", (done) => {
            const todoId = "3211afds";
            chai.request(app)
                .delete("/api/v1/todo/" + todoId)
                .end((err, response) => {
                    response.body.should.have.property("success").eq(false);
                    response.body.should.have.property("message").eq("Invalid Id");
                    done();
                });
        });

    });

});