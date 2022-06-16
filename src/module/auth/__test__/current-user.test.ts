import request from 'supertest'
import { app } from '../../../app'


it('response a data if authrized', async ()=>{
    const token = await global.signin()

    const response = await request(app)
        .get('/api/users/currentuser')
        .set('Cookie', token)
        .send()
        .expect(200)
    expect(response.body.currentUser.email).toEqual('test@test.com')
})

it('response a null if not authorized', async ()=>{

    const response = await request(app)
        .get('/api/users/currentuser')
        .send()
        .expect(200)
    expect(response.body.currentUser).toEqual(null)
})