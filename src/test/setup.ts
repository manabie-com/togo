import { MongoMemoryServer } from 'mongodb-memory-server'
import mongoose from 'mongoose'
import request from 'supertest'
import { app } from '../app'

declare global{
    function signin(): Promise<string[]>
}

jest.setTimeout(1000000)


let mongo:any
beforeAll(async ()=>{
    process.env.JWT_KEY = 'thuan'
    process.env.TLS_EJECT_UNAUTHORIZED = '0'
    mongo = new MongoMemoryServer()
    const mongoUri = await mongo.getUri()
    await mongoose.connect(mongoUri,{
        useUnifiedTopology: true,
        useNewUrlParser: true
    })
})

beforeEach(async ()=>{
    const collections = await mongoose.connection.db.collections()
    for(let collection of collections){
        await collection.deleteMany({})
    }
})

afterAll(async ()=>{
    await mongo.stop()
    await mongoose.connection.close()
})

global.signin = async () =>{
    const email = 'test@test.com'
    const password = 'password'
    const response = await request(app)
        .post('/api/users/signup')
        .send({
            email,
            password
        })
        .expect(201)

    const token = response.get('Set-Cookie')
    return token
}