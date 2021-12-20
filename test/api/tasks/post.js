// During the test the env variable is set to "test"
process.env.NODE_ENV = 'test';
const { TASK_LIMIT } = require('../../../config');

const expect = require('chai').expect;
const request = require('supertest');

const app = require('../../../app');
const { knex, db } = require('../../../db/knex');

describe('POST /tasks', () => {
    before((done) => {
        // Clear the test database
        const stmt = `
            DELETE FROM tasks;
            DELETE FROM user_tasks;
        `;
        db.exec(stmt, (err) => { if (err) throw err });
        done();
    });

    const validBody = {
        title: 'grab a coffee',
        detail: 'extra milk',
        reporter_id: 1,
        due_at: '2021-12-31 23:59'
    };
    it('... with valid body once', (done) => {
        request(app).post('/tasks').send(validBody)
            .then((res) => {
                const body = res.body;
                expect(body).to.contain.property('id');
                expect(body).to.contain.property('title');
                expect(body).to.contain.property('detail');
                expect(body).to.contain.property('reporter_id');
                expect(body).to.contain.property('assignee_id');
                done();
            })
            .catch((err) => done(err));
    })

    const missingReporterId = {
        title: 'OOO',
        detail: 'AAA'
    }
    it('... without property reporter_id', (done) => {
        request(app).post('/tasks').send(missingReporterId)
            .then((res) => {
                const body = res.body;
                expect(body.details[0].message).to.equal('"reporter_id" is required');
                done();
            })
            .catch((err) => done(err));
    })

    const newValidBody = {
        title: 'go to interview',
        detail: 'wear something comfortable',
        reporter_id: 2,
        due_at: '2021-12-31'
    };
    function makeTest(times) {
        var testDesc = times === TASK_LIMIT ? `... failed at ${times} time(s) a day` : `... ${times} time(s) a day`;
        it(testDesc, (done) => {
            request(app).post('/tasks').send(newValidBody)
                .then((res) => {
                    var { message } = res.body;
                    if (message) {
                        expect(message).to.equal(`Reached task count limit of ${TASK_LIMIT}`);
                        done();
                    } else {
                        done();
                    }
                })
                .catch((err) => done(err));
        })
    }
    for (let i = 1; i <= TASK_LIMIT + 1; i++) { makeTest(i) };
})