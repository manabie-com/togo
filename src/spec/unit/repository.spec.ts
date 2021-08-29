import * as chai from 'chai';
import 'mocha';
import { AppContainer } from "../../core";
import { REPOS, TaskRepository, UserRepository } from "../../repository";
import { StringHelper } from "../../shared";
import moment = require("moment");
import chaiHttp = require('chai-http');

chai.use(chaiHttp);
const should = chai.should();
const container = AppContainer.Load();

describe('[UNIT] Test Repository', () => {

    describe('UserRepository', () => {
        const userRepository = container.get<UserRepository>(REPOS.UserRepository);
        it('it should Fetch user by id', (done) => {
            userRepository.findOne({id: 'firstUser'}).then((user) => {
                should.exist(user);
                if (user) {
                    user.should.be.a('object');
                    user.should.have.property('id');
                    user.id.should.be.eq('firstUser');
                }
                done();
            }).catch(done);
        });
    });

    describe('TaskRepository', () => {
        const taskRepository = container.get<TaskRepository>(REPOS.TaskRepository);

        before((done) => {
            taskRepository.delete({}).then(res => done()).catch(done);
        });

        after((done) => {
            taskRepository.delete({}).then(res => done()).catch(done);
        });

        it('it should fetch list tasks', (done) => {
            taskRepository.find({}).then(tasks => {
                tasks.should.be.a('array');
                done();
            }).catch(done);
        });

        it('it should create new tasks', (done) => {
            const body: any = {
                id: StringHelper.generateUUID(),
                content: 'Lorem Ipsum',
                user_id: 'firstUser',
                created_date: moment().format('YYYY-MM-DD')
            };

            taskRepository.create(body).then(task => {
                task.should.be.a('object');
                task.should.have.property('content');
                task.content.should.be.contain('Lorem Ipsum');
                done();
            }).catch(done);
        });
    });

});
