// During the test the env variable is set to "test"
process.env.NODE_ENV = 'test';

const expect = require('chai').expect;
const request = require('supertest');

const app = require('../../../app');
const { knex, db } = require('../../../db/knex');

describe('GET /tasks', () => {
    before((done) => {
        // Clear the test database
        db.exec('DELETE FROM tasks', (err) => { if (err) throw err });
        done();
    });

    it('... with empty database', (done) => {
        request(app).get('/tasks')
            .then((res) => {
                const body = res.body;
                expect(body.length).to.equal(0);
                done();
            })
            .catch((err) => done(err));
    })

    const validBody = {
        title: 'grab a coffee',
        detail: 'extra milk',
        reporter_id: 3,
        due_at: '2021-12-31 23:59'
    };
    it('... with 1 task', (done) => {
        request(app).post('/tasks').send(validBody)
            .then((res) => {
                request(app).get('/tasks')
                    .then((res) => {
                        const body = res.body;
                        expect(body.length).to.equal(1);
                        done();
                    })
            })
            .catch((err) => done(err));
    })
})