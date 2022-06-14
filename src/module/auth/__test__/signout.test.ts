import request from 'supertest'
import { app } from '../../../app'


it('clear a cookie if singout successfull', async () => {
    const response = await request(app)
        .post('/api/users/signout')
        .send()
        .expect(200)
    expect(response.get('Set-Cookie')).toEqual(
        ['express:sess=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT; httponly']
    )
})