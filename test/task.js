process.env.NODE_ENV = 'test';

let chai = require('chai');
let chaiHttp = require('chai-http');
let server = require('../server');
let should = chai.should();

const taskUserModel = require('../models/task_user_model');

chai.use(chaiHttp);
describe('Tasks', () => {
    beforeEach((done) => {
        done();
    });

    describe('/POST tasks', () => {
        it('it should POST a task', (done) => {
            let task = {
                "user_id": 1,
                "task_name": 'AAA'
            }
            chai.request(server)
                .post('/tasks')
                .send(task)
                .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('date');
                    res.body.should.have.property('task_name').eql(task.task_name);
                    res.body.should.have.property('user_id').eql(task.user_id);
                    done();
                })
        });

        it('it should POST a task maximum limit of N tasks per user', (done) => {
            taskUserModel.save({
                "user_id": 1,
                "task_name": 'AAA',
                "date": "2022-2-21"
            });
            taskUserModel.save({
                "user_id": 1,
                "task_name": 'BBB',
                "date": "2022-2-21"
            });
            let task = {
                "user_id": 1,
                "task_name": 'CCC'
            }
            chai.request(server)
                .post('/tasks')
                .send(task)
                .end((err, res) => {
                    res.should.have.status(403);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql("You can not add task now");
                    done();
                })
        });

        it('it should not POST a task without task_name', (done) => {
            let task = {
                "user_id": 1
            }
            chai.request(server)
                .post('/tasks')
                .send(task)
                .end((err, res) => {
                    res.should.have.status(401);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql("invalid data");
                    done();
                });
        });

        it('it should not POST a task without user_id', (done) => {
            let task = {
                "task_name": 'AAA'
            }
            chai.request(server)
                .post('/tasks')
                .send(task)
                .end((err, res) => {
                    res.should.have.status(401);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql("invalid data");
                    done();
                });
        });
    });
});