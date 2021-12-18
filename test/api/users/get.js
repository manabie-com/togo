// During the test the env variable is set to "test"
process.env.NODE_ENV = 'test';

const expect = require('chai').expect;
const request = require('supertest');

const app = require('../../../app');
const { knex, db } = require('../../../db/knex');

describe('GET /users', () => {
    before((done) => {
        // Clear the test database
        db.exec('DELETE FROM users', (err) => { if (err) throw err });
        done();
    });

    it('... with empty database', (done) => {
        request(app).get('/users')
            .then((res) => {
                const body = res.body;
                expect(body.length).to.equal(0);
                done();
            })
            .catch((err) => done(err));
    })

    const validBody = {
        name: 'AAA',
        email: 'AAA@mail.com'
    };
    it('... with 1 user', (done) => {
        request(app).post('/users').send(validBody)
            .then((res) => {
                request(app).get('/users')
                    .then((res) => {
                        const body = res.body;
                        expect(body.length).to.equal(1);
                        done();
                    })
            })
            .catch((err) => done(err));
    })
})