import * as chai from 'chai';
import { request } from 'chai';
import 'mocha';
import { Server } from "../../index";
import { REPOS, TaskRepository } from "../../repository";
import { AppContainer } from "../../core";
import chaiHttp = require('chai-http');

chai.use(chaiHttp);
const expect = chai.expect;
const should = chai.should();
const app = Server.createServer();
const container = AppContainer.Load();

describe('[Integration] Test Request API', () => {

    const taskRepository = container.get<TaskRepository>(REPOS.TaskRepository);
    const userCredentials = {user_id: 'firstUser', password: 'example'};

    before((done) => {
        taskRepository.delete({}).then(res => done()).catch(done);
    });

    after((done) => {
        taskRepository.delete({}).then(res => done()).catch(done);
    });

    describe('[GET] /tasks', async () => {
        it('it should Login, then fetch list Tasks', (done) => {
            request(app).get('/login').query(userCredentials).end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('data');
                res.body.data.should.be.a('string');

                const token = res.body.data;
                request(app).get('/tasks').set('Authorization', token).end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('data');
                    res.body.data.should.be.a('array');
                    res.body.data.length.should.be.a('number');
                    done();
                });
            });
        });
    });

    describe('[POST] /tasks', async () => {
        let token: string = '';

        before((done) => {
            request(app).get('/login').query(userCredentials).end(function (err, res) {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('data');
                res.body.data.should.be.a('string');
                token = res.body.data;
                done();
            });
        });

        it('it should Create a FIRST Task', (done) => {
            const payload = {content: 'Lorem Ipsum #1'};
            request(app).post('/tasks').set('Authorization', token).send(payload).end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('data');
                res.body.data.should.be.a('object');
                res.body.data.content.should.be.contain(payload.content);
                done();
            });
        });

        it('it should Create a SECOND Task', (done) => {
            const payload = {content: 'Lorem Ipsum #2'};
            request(app).post('/tasks').set('Authorization', token).send(payload).end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('data');
                res.body.data.should.be.a('object');
                res.body.data.content.should.be.contain(payload.content);
                done();
            });
        });

        it('it should Create a THIRD Task', (done) => {
            const payload = {content: 'Lorem Ipsum #3'};
            request(app).post('/tasks').set('Authorization', token).send(payload).end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('data');
                res.body.data.should.be.a('object');
                res.body.data.content.should.be.contain(payload.content);
                done();
            });
        });

        it('it should Create a FORTH Task', (done) => {
            const payload = {content: 'Lorem Ipsum #4'};
            request(app).post('/tasks').set('Authorization', token).send(payload).end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('data');
                res.body.data.should.be.a('object');
                res.body.data.content.should.be.contain(payload.content);
                done();
            });
        });

        it('it should Create a FINAL Task', (done) => {
            const payload = {content: 'Lorem Ipsum #5'};
            request(app).post('/tasks').set('Authorization', token).send(payload).end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('data');
                res.body.data.should.be.a('object');
                res.body.data.content.should.be.contain(payload.content);
                done();
            });
        });

        it('it should Create more Task but limit and reject', (done) => {
            const payload = {content: 'Lorem Ipsum #6'};
            request(app).post('/tasks').set('Authorization', token).send(payload).end((err, res) => {
                res.should.have.status(400);
                res.body.should.be.a('object');
                res.body.should.have.property('error');
                res.body.error.should.be.contain('LIMIT_TODO_REACHED');
                done();
            });
        });
    });
});
