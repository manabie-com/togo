import request from 'supertest'
import { app } from '../../../app'

it('returns 401 if not authorized', async () => {
    await request(app)
    .get('/api/tasks')
    .send([
        {
            description: 'test1111',
            title: 'Test'
        },
        {
            description: 'test1111',
            title: 'Test'
        }
    ])
    .expect(401) 
})

it('returns 200 if success get list', async () => {
    const token = await global.signin()
    const response = await request(app)
        .get('/api/tasks')
        .set('Cookie', token )
        .expect(200)
})


it('check create task save and get list', async () => {
    const token = await global.signin()
    const responseGetListTaskAfterCreate = await request(app)
        .get('/api/tasks')
        .set('Cookie', token )
        .expect(200)
    const responseCreateTask = await request(app)
        .post('/api/tasks')
        .set('Cookie', token )
        .send([
            {
                description: 'test1111',
                title: 'Test'
            },
            {
                description: 'test1111',
                title: 'Test'
            }
        ])
        .expect(201)
    const responseGetListTaskBefore = await request(app)
        .get('/api/tasks')
        .set('Cookie', token )
        .expect(200)
    expect(responseGetListTaskAfterCreate.body.length).toBeLessThan(responseGetListTaskBefore.body.length)

})
