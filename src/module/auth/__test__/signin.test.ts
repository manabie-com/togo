import request from 'supertest'
import { app } from '../../../app'


it('returns 400 if invalid email and password', async () => {
    await request(app)
        .post('/api/users/signup')
        .send({
            email: 'test@test.com',
            password: 'password'
        })
        .expect(201)
    await request(app)
        .post('/api/users/signin')
        .send({
            email: 'test@test.com',
            password: '123321'
        })
        .expect(400)
    await request(app)
        .post('/api/users/signin')
        .send({
            email: 'test1@test.com',
            password: 'password'
        })
        .expect(400)
})

it('set a cookie if valid input', async () => {
    await request(app)
        .post('/api/users/signup')
        .send({
            email: 'test@test.com',
            password: 'password'
        })
        .expect(201)
    const response = await request(app)
        .post('/api/users/signin')
        .send({
            email: 'test@test.com',
            password: 'password'
        })
        .expect(200)

    expect(response.get('Set-Cookie')).toBeDefined()
})