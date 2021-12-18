// During the test the env variable is set to "test"
process.env.NODE_ENV = 'test';

const expect = require('chai').expect;
const request = require('supertest');

const app = require('../../../app');
const { knex, db } = require('../../../db/knex');

describe('POST /users', () => {
    before((done) => {
        // Clear the test database
        db.exec('DELETE FROM users', (err) => { if (err) throw err });
        done();
    });

    const validBody = {
        name: 'frank',
        email: 'frank@gmail.com'
    };
    it('... with valid body once', (done) => {
        request(app).post('/users').send(validBody)
            .then((res) => {
                const body = res.body;
                expect(body).to.contain.property('id');
                expect(body).to.contain.property('name');
                expect(body).to.contain.property('email');
                expect(body).to.contain.property('created_at');
                done();
            })
            .catch((err) => done(err));
    })
    
    it('... with the same email failed', (done) => {
        request(app).post('/users').send(validBody)
            .then((res) => {
                const body = res.body;
                expect(body.message).to.equal('email duplicated');
                done();
            })
            .catch((err) => done(err));
    })
})