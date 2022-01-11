const app = require('../../src/app');
const request = require('supertest');
const { before } = require('lodash');


describe("[INTEGRATION TEST]: HOME PAGE TEST.", () => {
    // Test base running
    beforeAll(async () => {
        console.log('This action running on before all.')
    });

    afterEach(async () => {
        console.log('This action running after each.')
    });

    afterAll(async () => {
        console.log('This action running after all.')
    });

    describe('Group test 01', () => {
        test('Homepage info private', async () => {
            const response = await request(app.callback()).get('/');
            expect(response.status).toBe(401);
        });

        test('Homepage info public', async () => {
            const response = await request(app.callback()).get('/api/public/home');
            expect(response.status).toBe(200);
        });
    })
})