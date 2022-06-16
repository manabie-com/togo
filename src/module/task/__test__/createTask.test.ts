import request from 'supertest'
import { app } from '../../../app'

it('returns 400 if invalid input', async () => {
    const token = await global.signin()
    await request(app)
        .post('/api/tasks')
        .set('Cookie', token)
        .send({
            description: 'test1111',
            title: 'Test'
        })
        .expect(400)  
    await request(app)
        .post('/api/tasks')
        .set('Cookie', token)
        .send([
            {
                description: 'test1111',
                title: 'Test'
            },
            {
                title: 'Test'
            }
        ])
        .expect(400) 
    await request(app)
        .post('/api/tasks')
        .set('Cookie', token)
        .send([
            {
                title: 'Test'
            },
            {
                description: 'test1111',
                title: 'Test'
            }
        ])
        .expect(400) 

})

it('returns 401 if not authorized', async () => {
    await request(app)
    .post('/api/tasks')
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

it('returns 201 if valid input', async () => {
    const token = await global.signin()
    const response = await request(app)
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
})