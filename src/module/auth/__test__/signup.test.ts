import request from 'supertest'
import {app} from '../../../app'


it('returns a 400 if invalid email and password', async ()=>{
    await request(app)
        .post('/api/users/signup')
        .send({
            password: 'password'
        })
        .expect(400)
    await request(app)
        .post('/api/users/signup')
        .send({
            email: '',
            password: 'password'
        })
        .expect(400)
    await request(app)
        .post('/api/users/signup')
        .send({
            email: 'test@test.com'
        })
        .expect(400)
    await request(app)
        .post('/api/users/signup')
        .send({
            email: 'test@test.com',
            password: ''
        })
        .expect(400)
})
it('disallows dupplicate email', async ()=>{
    await request(app)
        .post('/api/users/signup')
        .send({
            email: 'test@test.com',
            password: 'password'
        })
        .expect(201)
    await request(app)
        .post('/api/users/signup')
        .send({
            email: 'test@test.com',
            password: 'passwor1d'
        })
        .expect(400)
})

it('get a cookie if valid input', async ()=>{
    const response = await request(app)
        .post('/api/users/signup')
        .send({
            email: 'test@test.com',
            password: 'password'
        })
        .expect(201)
    expect(response.get('Set-Cookie')).toBeDefined()
})